package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/tkmagesh/Cisco-AdvGo-Jan-2025/01-recap/02-modules/utils"
)

func main() {
	color.Red("App executed")
	fmt.Println(utils.Add(100, 200))
	fmt.Println(utils.Subtract(100, 200))
}
