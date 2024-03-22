package main

import (
	"log"

	"gitea.bee.anarckk.me/anarckk/go_util/go_bit"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	_byte := go_bit.Int64ArrayToByteArray([]int64{123445, 9812358214})
	log.Println(_byte)
	i64Array := go_bit.ByteArrayToInt64Array(_byte)
	log.Println(i64Array)
}
