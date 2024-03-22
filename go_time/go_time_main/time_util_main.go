package main

import (
	"log"

	"gitea.bee.anarckk.me/anarckk/go_util/go_time"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
func test1() {
	log.Println(go_time.CurrentDatetimeStr())
	log.Println(go_time.CurrentDateStr())
	log.Println(go_time.CurrentTimeMillisStr())
	var str = go_time.FormatTime(go_time.ConvertUTCToLocal(go_time.GetNowUTC()))
	log.Println(str)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(go_time.CurrentTimeSecond())
}
