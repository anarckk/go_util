package go_os

import (
	"log"
	"testing"
)

func TestSimple(t *testing.T) {
	log.Println("isLinux", IsLinux())

	if IsWindows() {
		log.Println("当前是Windows环境")
	} else if IsLinux() {
		log.Println("当前是Linux环境")
	} else {
		log.Println("无法确定当前环境")
	}
}
