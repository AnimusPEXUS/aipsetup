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

	lst := []string{
		"iptables",
		"ipset",
		"conntrack-tools",
		"ulogd2",
		"nfacct",
		"arptables",
		"ebtables",

		"libmnl",
		"libnfnetlink",
		"libnetfilter_acct",
		"libnetfilter_conntrack",
		"libnetfilter_cttimeout",
		"libnetfilter_cthelper",
		"libnetfilter_queue",
		"libnetfilter_log",

		"libnftnl",
		"nftables",
		"nft-sync",

		"ulogd",
	}

	for _, i := range lst {
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
			"git://git.netfilter.org/" + i,
			i,
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
