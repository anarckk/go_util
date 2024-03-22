package go_string

import (
	"regexp"
	"strconv"
)

type ExcelUtil struct {
}

// ParseExcelCoordinate 解析excel的坐标值，拆开列和行
// coordinate A1
func (excelUtil *ExcelUtil) ParseExcelCoordinate(coordinate string) (column string, row int) {
	regex := regexp.MustCompile(`([A-Z]+)(\d+)`)
	matches := regex.FindStringSubmatch(coordinate)
	column = matches[1]
	row = parseInt(matches[2])
	return column, row
}

func parseInt(str string) int {
	var result int
	for _, char := range str {
		result = result*10 + int(char-'0')
	}
	return result
}

// IsExcelCoordinateValid 输入是否是合法的excel坐标值
// coordinate A1
func (excelUtil *ExcelUtil) IsExcelCoordinateValid(coordinate string) bool {
	// 使用正则表达式验证字符串格式是否正确
	match, _ := regexp.MatchString("^[A-Za-z]+\\d+$", coordinate)
	if !match {
		return false
	}

	// 解析列号和行号
	column := coordinate
	row := ""

	for i := 0; i < len(coordinate); i++ {
		if coordinate[i] >= '0' && coordinate[i] <= '9' {
			column = coordinate[:i]
			row = coordinate[i:]
			break
		}
	}

	// 检查列号是否合法
	if !excelUtil.isColumnValid(column) {
		return false
	}

	// 检查行号是否合法
	rowNumber, err := strconv.Atoi(row)
	if err != nil || rowNumber < 1 {
		return false
	}

	return true
}

func (excelUtil *ExcelUtil) isColumnValid(column string) bool {
	// 验证列号是否合法
	match, _ := regexp.MatchString("^[A-Za-z]+$", column)
	return match
}

// 计算列是第几列，从1开始， A 就是 1，Z是26, AA 就是27
func (excelUtil *ExcelUtil) ExcelColumn2Number(column string) int {
	result := 0
	for _, c := range column {
		result *= 26
		result += int(c) - 'A' + 1
	}
	return result
}

// 根据excel第n列计算它的列名
func (excelUtil *ExcelUtil) ExcelNumber2Column(n int) string {
	// A: 65, B: 66, ..., Z: 90
	const base = 26
	var result []byte
	for n > 0 {
		remainder := (n - 1) % base
		result = append([]byte{byte(65 + remainder)}, result...)
		n = (n - 1) / base
	}
	return string(result)
}
