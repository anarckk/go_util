/**
 * Author: anarckk anarckk@gmail.com
 * Date: 2023-05-04 13:49:21
 * LastEditTime: 2023-05-04 14:55:22
 * Description:
 *
 * Copyright (c) 2023 by anarckk, All Rights Reserved.
 */
package go_random

import (
	"fmt"
	"strconv"
	"testing"
)

func TestA(t *testing.T) {
	for i := 0; i < 1000; i++ {
		fmt.Print(strconv.Itoa(RandomInt2(0, 10)) + ", ")
		if (i+1)%50 == 0 {
			fmt.Println()
		}
	}
}
