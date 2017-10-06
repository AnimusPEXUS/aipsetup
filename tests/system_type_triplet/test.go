package main

import (
	"fmt"

	"github.com/AnimusPEXUS/aipsetup"
)

func main() {

	for ii, i := range []string{
		"x86_64-pc-linux-gnu",
		"i686-pc-linux-gnu",
		"x86_64-pc-windows",
		"i686-pc-windows",
	} {

		fmt.Printf("parsing test #%02d, subject: '%s'\n", ii, i)

		res, err := aipsetup.NewSystemTripletFromString(i)
		if err != nil {
			fmt.Println(" parsing error", err)
		} else {
			fmt.Printf(" CPU: %s, Company: %s, ", res.CPU, res.Company)
			if res.Kernel != "" {
				fmt.Printf("Kernel: %s, ", res.Kernel)
			}
			fmt.Printf("OS: %s\n", res.OS)
		}
		fmt.Println()
	}

}
