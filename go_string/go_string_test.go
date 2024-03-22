package go_string

import (
	"fmt"
	"testing"
)

func TestExcelColumn(t *testing.T) {
	excelUtil := &ExcelUtil{}
	for i := 1; i < 26*200+1; i++ {
		column := excelUtil.ExcelNumber2Column(i)
		i2 := excelUtil.ExcelColumn2Number(column)
		fmt.Print(column + ", ")
		if i%26 == 0 {
			fmt.Println()
		}
		if i != i2 {
			t.Errorf("%d 时不相等", i)
		}
	}
}
