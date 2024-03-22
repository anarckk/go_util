package go_code

// GetStrWidthSpace 获取字符串的宽度空间
// 一个中文占2个宽度空间，一个英文占1个字符空间
func GetStrWidthSpace(str string) int {
	var width = 0
	for _, s := range str {
		if len(string(s)) == 1 {
			width += 1
		} else {
			width += 2
		}
	}
	return width
}
