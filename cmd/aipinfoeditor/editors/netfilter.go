package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

func main() {

	infojson_dir := "../infojson"

	lst := [][]string{
		[]string{"iptables"},
		[]string{"ipset"},
		[]string{"conntrack-tools", "conntrack-tools"},
		[]string{"ulogd2", "ulogd"},
		[]string{"nfacct", "nfacct"},
		[]string{"arptables", "arptables"},
		[]string{"ebtables", "ebtables"},

		[]string{"libmnl", "libmnl"},
		[]string{"libnfnetlink", "libnfnetlink"},
		[]string{"libnetfilter_acct", "libnetfilter_acct"},
		[]string{"libnetfilter_conntrack", "libnetfilter_conntrack"},
		[]string{"libnetfilter_cttimeout", "libnetfilter_cttimeout"},
		[]string{"libnetfilter_cthelper", "libnetfilter_cthelper"},
		[]string{"libnetfilter_queue", "libnetfilter_queue"},
		[]string{"libnetfilter_log", "libnetfilter_log"},

		[]string{"libnftnl", "libnftnl"},
		[]string{"nftables"},
		[]string{"nft-sync"},

		[]string{"ulogd"},
	}

	for _, items := range lst {
		i := items[0]
		tag_prefix := "v"
		if len(items) > 1 {
			tag_prefix = items[1]
		}

		file_name := path.Join(infojson_dir, i+".json")
		_, err := os.Stat(file_name)
		exists := true
		if err != nil {
			if !os.IsNotExist(err) {
				panic(err)
			} else {
				exists = false
			}
		}

		info := new(basictypes.PackageInfo)

		if exists {
			data, err := ioutil.ReadFile(file_name)
			if err != nil {
				panic(err)
			}

			err = json.Unmarshal(data, info)
		}

		info.TarballName = i
		info.TarballFileNameParser = "std"
		info.HomePage = "https://git.netfilter.org/"
		info.TarballProvider = "srs"
		info.TarballProviderArguments = []string{
			"git",
			"git://git.netfilter.org/" + i,
			i,
		}
		switch tag_prefix {
		case "v":
		case "":
			info.TarballProviderArguments = append(info.TarballProviderArguments, "TagName:^$")
		default:
			info.TarballProviderArguments = append(info.TarballProviderArguments, "TagName:^"+tag_prefix+"$")
		}

		info.TarballVersionTool = "std"
		info.TarballProviderVersionSyncDepth = 3

		info.Category = "nettools"
		info.Groups = append(info.Groups, "netfilter")

		data, err := json.Marshal(info)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(file_name, data, 0700)
		if err != nil {
			panic(err)
		}

	}

}
