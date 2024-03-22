package beauty_print2

import (
	"fmt"
	"testing"
)

// go test --run TestBeautyPrint > a.txt
func TestBeautyPrint(t *testing.T) {
	path := "C:/data/syncthing/note_of_technology/"
	fmt.Println("test start.")
	file, err := GetFile(path)
	if err == nil {
		BeautyPrint(file, "", true)
	} else {
		fmt.Println(err)
	}
}
