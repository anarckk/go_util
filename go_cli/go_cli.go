package go_cli

import (
	"fmt"

	"gitea.bee.anarckk.me/anarckk/go_util/go_code"
)

// GetCliPrinter 获得一个命令行输出器，会重复在一行中输出字符串
func GetCliPrinter() func(string) {
	var _str string
	var clean = func() {
		for i := 0; i < len(_str); i++ {
			fmt.Print("\b")
		}
		for i := 0; i < go_code.GetStrWidthSpace(_str); i++ {
			fmt.Print(" ")
		}
		for i := 0; i < len(_str); i++ {
			fmt.Print("\b")
		}
	}
	return func(s string) {
		clean()
		_str = s
		fmt.Print(_str)
	}
}
