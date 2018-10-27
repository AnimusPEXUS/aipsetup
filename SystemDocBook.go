package aipsetup

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/AnimusPEXUS/utils/logger"
	"github.com/antchfx/xquery/xml"
)

const _ = `
OLD_DOCBOOK_INSTRUCTION

1. Get sources:
    aipsetup pkg-client-src get-lat docbook-sgml3
        will get docbk31.zip
    aipsetup pkg-client-src get-lat docbook-sgml4
        will get docbook-4.5.zip
    aipsetup pkg-client-src get-lat docbook-xml4
        will get docbook-xml-4.5.zip
    aipsetup pkg-client-src get-lat docbook-xsl
        will get docbook-xsl-[some version].tar[.some compressor]

    BIG FAT NOTE: YOU DO NOT NEED ALPHAS, BETAS AND RELISE CANDIDATES!!!
                  DO NOT PLAY WITH VERSIONS SUCH AS:
                  docbook-4.5b or docbook-4.5CR.

                  IF AIPSETUP DOWNLOADS SUCH FILES DO NOT TRUST IT AND
                  DOWNLOAD RIGHT FILES MANUALLY!!

2. Build files (root rights not required for this):
   aipsetup build full -d *

3. Install completed files (root rights required):
   aipsetup sys install *

4. Use command "aipsetup docbook install" as root

At this point docbook must be installed

NOTE: xmlcatalog command from vanilla libxml2-2.9.2.tar.gz has bug,
      which leads to incorrect work with xml catalog, which leads to
      damagede docbook installation

      following patch need to be appyed on libxml2-2.9.2 for xmlcatalog
      to work properly:

sed \
  -e /xmlInitializeCatalog/d \
  -e 's/((ent->checked =.*&&/(((ent->checked == 0) ||\
          ((ent->children == NULL) \&\& (ctxt->options \& XML_PARSE_NOENT))) \&\&/' \
  -i parser.c

    this patch is taken from here:
    http://www.linuxfromscratch.org/blfs/view/stable/general/libxml2.html
`

const DOCBOOK_INSTRUCTION = `
DocBook (re)installation instruction
-------------------------------------------------------------------------------
NOTE:
	  xmlcatalog command from vanilla libxml2-2.9.2.tar.gz has bug,
      which leads to incorrect work with xml catalog, which leads to
      damageded docbook installation

      following patch need to be appyed on libxml2-2.9.2 for xmlcatalog
      to work properly:

sed \
  -e /xmlInitializeCatalog/d \
  -e 's/((ent->checked =.*&&/(((ent->checked == 0) || \
   ((ent->children == NULL) \&\& (ctxt->options \& XML_PARSE_NOENT))) \&\&/' \
  -i parser.c

    this patch is taken from here:
    http://www.linuxfromscratch.org/blfs/view/stable/general/libxml2.html
-------------------------------------------------------------------------------
1. Get sources:
    aipsetup mbuild get-src docbook-sgml3 docbook-sgml4 docbook-xml4
		will get docbk31.zip, docbook-4.5.zip, docbook-xml-4.5.zip
		and docbook-xsl-[some version].tar[.some compressor]

    NOTE:
		you do not need alphas, betas and relise candidates!!!
		do not play with versions such as:
		docbook-4.5b or docbook-4.5CR.

		if aipsetup downloads such files do not trust it and
		download right files manually!!

2. Build files (root rights not required for this):
   aipsetup mbuild run

3. Install completed files (root rights required):
   aipsetup sys install-asps 01.asps/*

4. Use command "aipsetup sys-docbook install" as root

At this point docbook must be ready to use.
`

type InstallDockBookSettings struct {
	BaseDir          string
	SuperCatalogSGML string
	SuperCatalogXML  string
	SysSGMLDir       string
	SysXMLDir        string
	XMLCatalog       string
	Log              *logger.Logger
}

func (self *InstallDockBookSettings) SetDefaults(host_triplet string) {
	self.BaseDir = "/"
	self.SuperCatalogSGML = "/etc/sgml/sgml-docbook.cat"
	self.SuperCatalogXML = "/etc/xml/docbook"
	self.SysSGMLDir = path.Join("/multihost", host_triplet, "/share/sgml/docbook")
	self.SysXMLDir = path.Join("/multihost", host_triplet, "/share/xml/docbook")
	self.XMLCatalog = "/etc/xml/catalog"
}

type DocBookCtl struct {
	settings *InstallDockBookSettings
}

func NewDocBookCtl(settings *InstallDockBookSettings) *DocBookCtl {
	ret := &DocBookCtl{settings: settings}
	return ret
}

func (self *DocBookCtl) LogI(txt string) {
	if self.settings.Log != nil {
		self.settings.Log.Info(txt)
	}
}

func (self *DocBookCtl) LogE(txt string) {
	if self.settings.Log != nil {
		self.settings.Log.Error(txt)
	}
}

func (self *DocBookCtl) InstallDockBook() error {

	//	BD_SuperCatalogSGML :=
	//		path.Join(self.settings.BaseDir, self.settings.SuperCatalogSGML)

	//	BD_SuperCatalogXML :=
	//		path.Join(self.settings.BaseDir, self.settings.SuperCatalogXML)

	BD_SysSGMLDir :=
		path.Join(self.settings.BaseDir, self.settings.SysSGMLDir)

	BD_SysXMLDir :=
		path.Join(self.settings.BaseDir, self.settings.SysXMLDir)

	BD_XMLCatalog :=
		path.Join(self.settings.BaseDir, self.settings.XMLCatalog)

	sgml_dirs := []string{
		"dockbook-3.1",
		"dockbook-4.5",
	}

	xml_dirs := []string{
		//		"dockbook-xml-4.5",
	}

	xsl_dirs := []string{
		//		"dockbook-xml-4.5",
	}

	for _, i := range sgml_dirs {
		p := path.Join(BD_SysSGMLDir, i)
		s, err := os.Stat(p)
		if err != nil {
			return errors.New(p + " not found")
		}
		if !s.IsDir() {
			return errors.New(p + " not dir")
		}
	}

	{
		files, err := ioutil.ReadDir(BD_SysXMLDir)
		if err != nil {
			return err
		}

		for _, i := range files {
			if s, err := regexp.MatchString(`docbook-xml-(\d+\.?)+`, i.Name()); err != nil {
				return err
			} else {
				if s {
					xml_dirs = append(xml_dirs, i.Name())
				}
			}
		}
	}

	{
		files, err := ioutil.ReadDir(BD_SysXMLDir)
		if err != nil {
			return err
		}

		for _, i := range files {
			if s, err := regexp.MatchString(`docbook-xsl-(\d+\.?)+`, i.Name()); err != nil {
				return err
			} else {
				if s {
					xsl_dirs = append(xsl_dirs, i.Name())
				}
			}
		}
	}

	if len(xml_dirs) != 1 {
		return errors.New("must be exactly one xml directory inside of " + BD_SysXMLDir)
	}

	if len(xsl_dirs) != 1 {
		return errors.New("must be exactly one xsl directory inside of " + BD_SysXMLDir)
	}

	for _, i := range sgml_dirs {
		target_dir := path.Join(self.settings.SysSGMLDir, i)

		self.ImportToSuperDocBookCatalog(
			target_dir,
			self.settings.BaseDir,
			self.settings.SuperCatalogSGML,
			self.settings.SuperCatalogXML,
		)

		if strings.HasSuffix(i, "4.5") {
			self.MakeNewDocbook_4_5_LookLikeOld(
				self.settings.BaseDir,
				target_dir,
			)
		}

		if strings.HasSuffix(i, "3.1") {
			self.MakeNewDocbook_3_1_LookLikeOld(
				self.settings.BaseDir,
				target_dir,
			)
		}

	}

	for _, i := range xml_dirs {
		target_dir := path.Join(self.settings.SysXMLDir, i)

		self.ImportToSuperDocBookCatalog(
			target_dir,
			self.settings.BaseDir,
			self.settings.SuperCatalogSGML,
			self.settings.SuperCatalogXML,
		)

		self.MakeNewDockBookXMLLookLikeOld(
			self.settings.BaseDir,
			target_dir,
			self.settings.SuperCatalogXML,
			self.settings.XMLCatalog,
		)

	}

	for _, i := range xsl_dirs {
		target_dir := path.Join(self.settings.SysXMLDir, i)

		self.ImportXSLToXMLCatalog(
			target_dir,
			self.settings.BaseDir,
			true, // TODO: this value is under question
			self.settings.XMLCatalog,
		)

	}

	self.ImportDocBookToCatalog(BD_XMLCatalog)

	return nil
}

func (self *DocBookCtl) ImportToSuperDocBookCatalog(
	target_dir string,
	base_dir string,
	super_catalog_sgml string,
	super_catalog_xml string,
) error {

	target_dir_fd := path.Join(base_dir, target_dir)

	files, err := ioutil.ReadDir(target_dir_fd)
	if err != nil {
		return err
	}

	for _, i := range files {

		if i.Name() == "docbook.cat" {
			c := exec.Command(
				// TODO: no guessing, use 'which'
				"xmlcatalog",
				"--sgml",
				"--noout",
				"--add",
				path.Join(target_dir_fd, "docbook.cat"),
				super_catalog_sgml,
			)
			err := c.Run()
			if err != nil {
				return err
			}
		}

		if i.Name() == "catalog.xml" {

			target_catalog_xml := path.Join(target_dir_fd, "catalog.xml")

			err := self.ImportCatalogXMLToSuperDocBookCatalog(
				target_catalog_xml,
				base_dir,
				super_catalog_xml,
			)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func (self *DocBookCtl) MakeNewDocbook_4_5_LookLikeOld(
	base_dir string,
	target_dir string,
) error {
	return self.MakeNewDocbook_x_x_LookLikeOld(base_dir, target_dir, "45")
}

func (self *DocBookCtl) MakeNewDocbook_3_1_LookLikeOld(
	base_dir string,
	target_dir string,
) error {
	return self.MakeNewDocbook_x_x_LookLikeOld(base_dir, target_dir, "31")
}

func (self *DocBookCtl) MakeNewDocbook_x_x_LookLikeOld(
	base_dir string,
	installed_docbook_xml_dir string,
	ver string,
) error {

	installed_docbook_xml_dir_fn :=
		path.Join(
			base_dir,
			installed_docbook_xml_dir,
		)

	catalog_fn := path.Join(
		installed_docbook_xml_dir_fn,
		"docbook.cat",
	)

	data, err := ioutil.ReadFile(catalog_fn)
	if err != nil {
		return err
	}

	data_str := string(data)

	data_str_lines := strings.Split(data_str, "\n")

	if ver == "31" {
		//    logging.info("Adding support for older docbook-4.* versions")

		for _, i := range []string{"3.0"} {
			//		logging.info("    {}".format(i))

			new_line :=
				`UBLIC "-//Davenport//DTD DocBook V` + i + `//EN" "docbook.dtd"` //+"\n"

			for _, j := range data_str_lines {
				if new_line == j {
					continue
				}
			}

			data_str_lines = append(data_str_lines, new_line)
		}
	} else if ver == "45" {

		//    logging.info("Adding support for older docbook-4.* versions")

		for _, i := range []string{"4.4", "4.3", "4.2", "4.1", "4.0"} {
			//		logging.info("    {}".format(i))

			new_line :=
				`PUBLIC "-//OASIS//DTD DocBook V` + i + `//EN" "docbook.dtd"` //+"\n"

			for _, j := range data_str_lines {
				if new_line == j {
					continue
				}
			}

			data_str_lines = append(data_str_lines, new_line)
		}

	} else {
		panic("invalid ver value")
	}

	err = ioutil.WriteFile(catalog_fn, []byte(strings.Join(data_str_lines, "\n")), 0700)
	if err != nil {
		return err
	}

	return nil
}

func (self *DocBookCtl) MakeNewDockBookXMLLookLikeOld(
	base_dir string,
	installed_docbook_xml_dir string,
	super_catalog_xml string,
	xml_catalog string,
) error {

	super_catalog_xml_fn := path.Join(
		base_dir,
		super_catalog_xml,
	)

	//    logging.info("Adding support for older docbook-xml versions")
	//    logging.info("    ({})".format(super_catalog_xml_fn))

	for _, i := range []string{"4.1.2", "4.2", "4.3", "4.4"} {

		//        logging.info("    {}".format(i))

		c := exec.Command(
			"xmlcatalog", "--noout", "--add", "public",
			`-//OASIS//DTD DocBook XML V`+i+`//EN`,
			"http://www.oasis-open.org/docbook/xml/"+i+"/docbookx.dtd",
			super_catalog_xml_fn,
		)

		err := c.Run()
		if err != nil {
			return err
		}

		c = exec.Command(
			"xmlcatalog", "--noout", "--add", "rewriteSystem",
			"http://www.oasis-open.org/docbook/xml/"+i,
			"file://"+installed_docbook_xml_dir,
			super_catalog_xml_fn,
		)

		err = c.Run()
		if err != nil {
			return err
		}

		c = exec.Command(
			"xmlcatalog", "--noout", "--add", "rewriteURI",
			"http://www.oasis-open.org/docbook/xml/"+i,
			"file://"+installed_docbook_xml_dir,
			super_catalog_xml_fn,
		)

		err = c.Run()
		if err != nil {
			return err
		}

		c = exec.Command(
			"xmlcatalog", "--noout", "--add", "delegateSystem",
			"http://www.oasis-open.org/docbook/xml/"+i,
			"file://"+super_catalog_xml,
			super_catalog_xml_fn,
		)

		err = c.Run()
		if err != nil {
			return err
		}

		c = exec.Command(
			"xmlcatalog", "--noout", "--add", "delegateURI",
			"http://www.oasis-open.org/docbook/xml/"+i,
			"file://"+super_catalog_xml,
			super_catalog_xml_fn,
		)

		err = c.Run()
		if err != nil {
			return err
		}

	}

	return nil
}

func (self *DocBookCtl) ImportXSLToXMLCatalog(
	target_xsl_dir string,
	base_dir string,
	current bool,
	xml_catalog string,
) error {

	//	target_xsl_dir_fn := path.Join(base_dir, target_xsl_dir)
	target_xsl_dir_fn_no_base := target_xsl_dir // TODO: ???

	xml_catalog_fn := path.Join(base_dir, xml_catalog)

	bn := path.Base(target_xsl_dir)

	version := strings.Replace(bn, "docbook-xsl-", "", -1)

	//    logging.info("Importing version: {}".format(version))

	c := exec.Command(
		"xmlcatalog", "--noout", "--add", "rewriteSystem",
		"http://docbook.sourceforge.net/release/xsl/"+version,
		target_xsl_dir_fn_no_base,
		xml_catalog_fn,
	)

	err := c.Run()
	if err != nil {
		return err
	}

	c = exec.Command(
		"xmlcatalog", "--noout", "--add", "rewriteURI",
		"http://docbook.sourceforge.net/release/xsl/"+version,
		target_xsl_dir_fn_no_base,
		xml_catalog_fn,
	)

	err = c.Run()
	if err != nil {
		return err
	}

	if current {

		c = exec.Command(
			"xmlcatalog", "--noout", "--add", "rewriteSystem",
			"http://docbook.sourceforge.net/release/xsl/current",
			target_xsl_dir_fn_no_base,
			xml_catalog_fn,
		)

		err = c.Run()
		if err != nil {
			return err
		}

		c = exec.Command(
			"xmlcatalog", "--noout", "--add", "rewriteURI",
			"http://docbook.sourceforge.net/release/xsl/current",
			target_xsl_dir_fn_no_base,
			xml_catalog_fn,
		)

		err = c.Run()
		if err != nil {
			return err
		}

	}
	return nil
}

func (self *DocBookCtl) ImportDocBookToCatalog(
	base_dir_etc_xml_catalog string,
) error {

	for _, each := range [][]string{
		[]string{
			"xmlcatalog", "--noout", "--add", "delegatePublic",
			"-//OASIS//ENTITIES DocBook XML",
			"file:///etc/xml/docbook", // TODO: no hardcode
			base_dir_etc_xml_catalog,
		},
		[]string{
			"xmlcatalog", "--noout", "--add", "delegatePublic",
			"-//OASIS//DTD DocBook XML",
			"file:///etc/xml/docbook", // TODO: no hardcode
			base_dir_etc_xml_catalog,
		},
		[]string{
			"xmlcatalog", "--noout", "--add", "delegateSystem",
			"http://www.oasis-open.org/docbook/",
			"file:///etc/xml/docbook", // TODO: no hardcode
			base_dir_etc_xml_catalog,
		},
		[]string{
			"xmlcatalog", "--noout", "--add", "delegateURI",
			"http://www.oasis-open.org/docbook/",
			"file:///etc/xml/docbook", // TODO: no hardcode
			base_dir_etc_xml_catalog,
		},
	} {

		c := exec.Command(each[0], each[1:]...)
		err := c.Run()
		if err != nil {
			return err
		}

	}

	return nil
}

func (self *DocBookCtl) ImportCatalogXMLToSuperDocBookCatalog(
	target_catalog_xml string,
	base_dir string,
	super_docbook_catalog_xml string,
) error {

	target_catalog_xml_fn := path.Join(base_dir, target_catalog_xml)

	//	target_catalog_xml_dir := path.Dir(target_catalog_xml)
	//	target_catalog_xml_fn_dir := path.Dir(target_catalog_xml_fn)

	//    target_catalog_xml_fn_dir_virtual = target_catalog_xml_fn_dir
	//    target_catalog_xml_fn_dir_virtual = wayround_i2p.utils.path.remove_base(
	//        target_catalog_xml_fn_dir_virtual, base_dir
	//        )
	//	target_catalog_xml_fn_dir_virtual := target_catalog_xml_dir

	super_docbook_catalog_xml_fn := path.Join(
		base_dir,
		super_docbook_catalog_xml,
	)
	super_docbook_catalog_xml_fn_dir := path.Dir(
		super_docbook_catalog_xml_fn,
	)

	err := os.MkdirAll(super_docbook_catalog_xml_fn_dir, 0755)
	if err != nil {
		return err
	}

	target_catalog_xml_fn_data, err := ioutil.ReadFile(target_catalog_xml_fn)
	if err != nil {
		return err
	}

	br := bytes.NewReader(target_catalog_xml_fn_data)
	//	bufferedio

	//	var tmp_cat_lxml

	tmp_cat_lxml, err := xmlquery.Parse(br)
	if err != nil {
		return err
	}

	node := tmp_cat_lxml.FirstChild
	for node != nil {
		if node.Type == xmlquery.ElementNode {
			fmt.Println("node.Data", node.Data)
			os.Exit(666)
		}
		node = node.NextSibling
	}

	//    for i in tmp_cat_lxml.getroot():

	//        if type(i) == lxml.etree._Element:

	//            qname = lxml.etree.QName(i.tag)

	//            tag = qname.localname

	//            src_uri = i.get('uri')

	//            if src_uri:

	//                dst_uri = ''

	//                if (src_uri.startswith('http://')
	//                        or src_uri.startswith('https://')
	//                        or src_uri.startswith('file://')):

	//                    dst_uri = src_uri

	//                else:

	//                    dst_uri = wayround_i2p.utils.path.join(
	//                        '/', target_catalog_xml_fn_dir_virtual, src_uri
	//                        )

	//                logging.info("    adding {}".format(i.get(tag + 'Id')))

	//                cmd = [
	//                    'xmlcatalog', '--noout', '--add',
	//                    tag,
	//                    i.get(tag + 'Id'),
	//                    'file://{}'.format(dst_uri),
	//                    super_docbook_catalog_xml_fn,
	//                    #i.get(tag + 'Id')
	//                    ]

	//                p = subprocess.Popen(cmd)

	//                p.wait()

	return nil
}
