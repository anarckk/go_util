package main

import (
	"fmt"
	"time"

	"gitea.bee.anarckk.me/anarckk/go_util/go_cli"
)

func main() {
	printer := go_cli.GetCliPrinter()
	printer("我喜欢用C语言写程序")
	time.Sleep(time.Second * 1)
	printer("也喜欢用go写")
	time.Sleep(time.Second * 1)
	printer("真的吗")
	time.Sleep(time.Second * 1)
	printer("还是擦吧")
	time.Sleep(time.Second * 1)
	printer("已擦除")
	fmt.Println()
}
