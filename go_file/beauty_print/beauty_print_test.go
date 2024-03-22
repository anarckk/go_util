/**
 * Author: anarckk anarckk@gamil.com
 * Date: 2023-06-26 10:37:52
 * LastEditTime: 2023-06-26 10:40:43
 * Description:
 *
 * Copyright (c) 2023 by anarckk, All Rights Reserved.
 */
package beauty_print

import (
	"testing"
)

func Test1(t *testing.T) {
	var file *MyFile
	var err error
	if file, err = GetFile("C:\\Users\\anarckk\\Downloads\\"); err != nil {
		panic(err)
	}
	BeautyPrint(file, "", true)
}
