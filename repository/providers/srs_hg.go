package providers

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/utils/tags"
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
	"github.com/AnimusPEXUS/utils/tarballversion"
	"github.com/AnimusPEXUS/utils/tarballversion/versioncomparators"
)

var HG_TAG = regexp.MustCompile(`(.*?)\s+(\d+\:[0-9a-f]+)$`)

type SRSHg struct {
	srs *ProviderSRS
}

func NewSRSHg(srs *ProviderSRS) *SRSHg {
	self := new(SRSHg)
	self.srs = srs
	return self
}

func (self *SRSHg) GetAndUpdate(
	hg_dir string,
	hg_source_url string,
) error {

	hg_subdir := path.Join(hg_dir, ".hg")

	new_download_mode := false

	if _, err := os.Stat(hg_subdir); err == nil {
		new_download_mode = false
	} else {
		new_download_mode = true
	}

	var err error

	if new_download_mode {
		err = os.RemoveAll(hg_dir)
		if err != nil {
			return err
		}

		err = os.MkdirAll(hg_dir, 0700)
		if err != nil {
			return err
		}

		self.srs.log.Info(fmt.Sprintf("getting %s", hg_source_url))

		c := exec.Command("hg", "clone", hg_source_url, ".")
		c.Dir = hg_dir
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		err = c.Run()
		if err != nil {
			return err
		}

	} else {

		self.srs.log.Info(fmt.Sprintf("updating %s", hg_source_url))

		c := exec.Command("hg", "checkout", "-C")
		c.Dir = hg_dir
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		err = c.Run()
		if err != nil {
			return err
		}

	}

	return nil
}

func (self *SRSHg) MakeTarballs(
	hg_dir string,
	output_dir string,
	info *basictypes.PackageInfo,
) error {

	t := tags.New(info.TarballProviderArguments)

	{
		tp := self.srs.repo.GetPackageTarballsPath(self.srs.pkg_name)

		err := os.MkdirAll(tp, 0700)
		if err != nil {
			return err
		}
	}

	TagParser := self.srs.pkg_info.TarballFileNameParser
	if tt, ok := t.GetSingle("TagParser", true); ok {
		TagParser = tt
	}

	TagName := "^v"
	if tt, ok := t.GetSingle("TagName", true); ok {
		TagName = tt
	}

	TagStatus := "^$"
	if tt, ok := t.GetSingle("TagStatus", true); ok {
		TagStatus = tt
	}

	TagComparator := self.srs.pkg_info.TarballVersionComparator
	if tt, ok := t.GetSingle("TagComparator", true); ok {
		TagComparator = tt
	}

	TagTarballRenderer := TagParser
	if tt, ok := t.GetSingle("TagTarballRenderer", true); ok {
		TagTarballRenderer = tt
	}

	TagFilters, TagFiltersUse := t.GetSingle("TagFilters", true)

	// TODO: do srs SyncDepth should also be moved from InfoDB to srs args?
	truncate_versions := info.TarballProviderVersionSyncDepth
	if truncate_versions == 0 {
		truncate_versions = 3
	}

	parser, err := tarballnameparsers.Get(TagParser)
	if err != nil {
		return err
	}

	comparator, err := versioncomparators.Get(TagComparator)
	if err != nil {
		return err
	}

	stability_classifier, err := tarballstabilityclassification.Get(
		self.srs.pkg_info.TarballStabilityClassifier,
	)
	if err != nil {
		return err
	}

	acceptable_tags := make([]string, 0)

	{
		var tags []string

		{
			b := &bytes.Buffer{}
			c := exec.Command("hg", "tags")
			c.Dir = hg_dir
			c.Stdout = b
			err := c.Run()
			if err != nil {
				return err
			}
			tags = strings.Split(b.String(), "\n")

			{
				tags2 := make([]string, 0)
				for _, i := range tags {
					if !HG_TAG.MatchString(i) {
						continue
					}
					r := HG_TAG.FindStringSubmatch(i)
					tags2 = append(tags2, r[1])
				}

				tags = tags2
			}

		}

		for _, i := range tags {
			parse_res, err := parser.Parse(i)
			if err != nil {
				continue
			}

			matched, err := regexp.MatchString(TagName, parse_res.Name)
			if err != nil {
				return err
			}

			if !matched {
				continue
			}

			matched, err = regexp.MatchString(
				TagStatus,
				parse_res.Status.DirtyString(),
			)
			if err != nil {
				return err
			}

			if !matched {
				continue
			}

			if TagFiltersUse {
				info := self.srs.pkg_info
				switch TagFilters {
				case "+":
					info = self.srs.pkg_info
				default:
					var err error
					info, err = pkginfodb.Get(TagFilters)
					if err != nil {
						return errors.New("can't get named info filters for srs")
					}
				}
				fres, err := pkginfodb.ApplyInfoFilter(
					info,
					[]string{i},
				)
				if err != nil {
					return err
				}

				if len(fres) != 1 {
					continue
				}

			}

			if ok, err := stability_classifier.IsStable(parse_res); err != nil {
				return err
			} else {
				if !ok {
					continue
				}
			}

			acceptable_tags = append(acceptable_tags, i)
		}
	}

	{

		version_tree, err := tarballversion.NewVersionTree(
			TagName,
			parser,
			comparator,
		)
		if err != nil {
			return err
		}

		for _, i := range acceptable_tags {
			b := path.Base(i)

			err = version_tree.Add(b)
			if err != nil {
				return err
			}
		}

		depth := self.srs.pkg_info.TarballProviderVersionSyncDepth
		if depth == 0 {
			depth = 3
		}

		self.srs.log.Info("-----------------")
		self.srs.log.Info("tags before versioned truncation")

		res, err := version_tree.Basenames([]string{""})
		if err != nil {
			return err
		}
		for _, i := range res {
			self.srs.log.Info(fmt.Sprintf("  %s", i))
		}

		version_tree.TruncateByVersionDepth(nil, depth)

		self.srs.log.Info("-----------------")
		self.srs.log.Info("tag versioned truncation result")

		res, err = version_tree.Basenames([]string{""})
		if err != nil {
			return err
		}
		for _, i := range res {
			self.srs.log.Info(fmt.Sprintf("  %s", i))
		}

		err = comparator.SortStrings(res, parser)
		if err != nil {
			return err
		}

		self.srs.log.Info("-----------------")
		self.srs.log.Info("sorted by version")

		for _, i := range res {
			self.srs.log.Info(fmt.Sprintf("  %s", i))
		}

		{
			len_res := len(res)
			t := make([]string, len_res)
			for i := range res {
				t[i] = res[len_res-i-1]
			}
			res = t
		}

		self.srs.log.Info("-----------------")
		self.srs.log.Info("to archive")

		for _, i := range res {
			self.srs.log.Info(fmt.Sprintf("  %s", i))
		}

		acceptable_tags = res

	}

	was_errors := false
	downloaded_files := make([]string, 0)

	{
		name_renderer, err := tarballnameparsers.Get(TagTarballRenderer)
		if err != nil {
			return err
		}

		self.srs.log.Info("-----------------")
		self.srs.log.Info("archiving")

		for _, i := range acceptable_tags {

			i_parsed, err := parser.Parse(i)
			if err != nil {
				return err
			}

			tag_filename_noext, err := name_renderer.Render(&tarballname.ParsedTarballName{
				Name:    info.TarballName,
				Version: i_parsed.Version,
				Status:  i_parsed.Status,
			})
			if err != nil {
				return err
			}

			tag_filename := tag_filename_noext + ".tar.bz2"

			tag_filename_done := self.srs.repo.GetTarballDoneFilePath(
				self.srs.pkg_name,
				tag_filename,
			)

			if _, err := os.Stat(tag_filename_done); err == nil {
				downloaded_files = append(downloaded_files, tag_filename)
				continue
			}

			target_file := self.srs.repo.GetTarballFilePath(self.srs.pkg_name, tag_filename)

			self.srs.log.Info(fmt.Sprintf("  %s", tag_filename))

			c := exec.Command(
				"hg",
				"archive",
				"-t", "tbz2",
				"-r", i,
				fmt.Sprintf("--prefix=%s/", tag_filename_noext),
				target_file,
			)
			c.Dir = hg_dir
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr

			if c.Run() != nil {
				was_errors = true
			} else {
				downloaded_files = append(downloaded_files, tag_filename)
				if f, err := os.Create(tag_filename_done); err != nil {
					return err
				} else {
					f.Close()
				}
			}

		}
	}

	if was_errors {
		return errors.New("there was errors making tarballs")
	}

	// TODO: maybe rest of this func shold be moved to PerformUpdate

	lst, err := self.srs.repo.PrepareTarballCleanupListing(self.srs.pkg_name, downloaded_files)
	if err != nil {
		return err
	}

	self.srs.log.Info("-----------------")
	self.srs.log.Info("to delete")

	for _, i := range lst {
		self.srs.log.Info(fmt.Sprintf("  %s", i))
	}

	err = self.srs.repo.DeleteTarballFiles(self.srs.pkg_name, lst)
	if err != nil {
		return err
	}

	return nil
}
