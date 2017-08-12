package main

import (
	"fmt"

	"github.com/AnimusPEXUS/aipsetup"
)

func main() {
	sys := aipsetup.NewSystem("/")
	fmt.Println(sys.Root())
	fmt.Println(sys.Host())
	fmt.Println(sys.Archs())
}
