package main

import (
	"fmt"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl/rtl"
)

func main() {
	fmt.Println(rtl.SysLocale)
}
