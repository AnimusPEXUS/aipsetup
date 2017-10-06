package main

import (
	"fmt"

	"github.com/AnimusPEXUS/aipsetup"
)

var EXAMPLES []string = []string{
	"(font-isas-misc)-(1.0.3)-()-(20150903.045208.0791884)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(font-jis-misc)-(1.0.3)-()-(20150903.045215.0840178)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(font-micro-misc)-(1.0.3)-()-(20150903.045221.0561060)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(font-misc-cyrillic)-(1.0.3)-()-(20150903.045227.0520193)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(font-misc-ethiopic)-(1.0.3)-()-(20150903.045233.0309248)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(font-misc-meltho)-(1.0.3)-()-(20150903.045241.0939742)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(font-misc-misc)-(1.1.2)-()-(20150903.045835.0602271)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(font-mutt-misc)-(1.0.3)-()-(20150903.045252.0924733)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(font-schumacher-misc)-(1.1.2)-()-(20150903.045842.0948272)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(font-screen-cyrillic)-(1.0.4)-()-(20150903.045301.0822368)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(font-sony-misc)-(1.0.3)-()-(20150903.045307.0251229)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(fontsproto)-(2.1.3)-()-(20160503.010231.0183374)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(fontsproto)-(2.1.3)-()-(20160503.010236.0319396)-(x86_64-pc-linux-gnu)-(i686-pc-linux-gnu).xz",
	"(font-sun-misc)-(1.0.3)-()-(20150903.045313.0010382)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(fonttosfnt)-(1.0.4)-()-(20160503.013831.0944624)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(fonttosfnt)-(1.0.4)-()-(20160503.013837.0116282)-(x86_64-pc-linux-gnu)-(i686-pc-linux-gnu).xz",
	"(font-util)-(1.3.1)-()-(20150903.045322.0015648)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(font-winitzki-cyrillic)-(1.0.3)-()-(20150903.045327.0321006)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(font-xfree86-type1)-(1.0.4)-()-(20150903.045331.0883884)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(foomatic-db-engine)-(4.0.12)-()-(20160702.021348.0585480)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(foomatic-db-engine)-(4.0.12)-()-(20160702.021355.0492161)-(x86_64-pc-linux-gnu)-(i686-pc-linux-gnu).xz",
	"(foomatic-filters)-(4.0.17)-()-(20160702.021404.0579351)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(foomatic-filters)-(4.0.17)-()-(20160702.021413.0598832)-(x86_64-pc-linux-gnu)-(i686-pc-linux-gnu).xz",
	"(fpc)-(3.0.0)-(source)-(20160528.210816.0757124)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(fpc)-(3.0.0)-(source)-(20160528.211140.0898816)-(x86_64-pc-linux-gnu)-(i686-pc-linux-gnu).xz",
	"(freealut)-(1.1.0)-()-(20160702.201144.0496965)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
	"(freealut)-(1.1.0)-()-(20160702.201157.0463384)-(x86_64-pc-linux-gnu)-(i686-pc-linux-gnu).xz",
	"(freeglut)-(3.0.0)-()-(20161128.200001.0405316)-(x86_64-pc-linux-gnu)-(x86_64-pc-linux-gnu).xz",
}

func main() {
	for ii, i := range EXAMPLES {
		v, err := aipsetup.NewASPNameFromString(i)
		fmt.Printf("Example #%02d:\n", ii)
		if err != nil {
			fmt.Println("error", err)
		} else {
			fmt.Printf("%s\n", i)
			fmt.Printf("%s\n", v.String())
			t, err := v.TimeStampTime()
			fmt.Printf("%v, %v, %v, %v\n", v.Name, v.TimeStamp, t, err)
		}
		fmt.Println()
	}
}
