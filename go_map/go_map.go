package go_map

import "strings"

// 用逗号把字符串数组组合起来
func ComposeStrArray(arr []string) string {
	return strings.Join(arr, ",")
}

// 用逗号把字符串数组组合起来
func ComposeArrayT[T any](arr []T, tr func(T) string) string {
	result := ""
	_len := len(arr)
	for i, a := range arr {
		result += tr(a)
		if i != _len - 1 {
			result += ", "
		}
	}
	return result
}
