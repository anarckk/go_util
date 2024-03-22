/**
 * Author: anarckk anarckk@gmail.com
 * Date: 2023-05-01 09:57:22
 * LastEditTime: 2023-05-01 09:57:47
 * Description:
 *
 * Copyright (c) 2023 by anarckk, All Rights Reserved.
 */
package go_string

import (
	"os"
	"strings"
	"time"

	"gitea.bee.anarckk.me/anarckk/go_util/go_random"
	uuid "github.com/satori/go.uuid"
)

// HasPrefix 判断是否有指定的字符串前缀
//
//	@param s
//	@param prefix
//	@return bool
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// GetUUID 返回一个uuid
// 示例：4642015db2fa4aff85b0990cd6e69251
func GetUUID() string {
	u2 := uuid.NewV4()
	return strings.ReplaceAll(u2.String(), "-", "")
}

// GetRandomStr 根据指定的chars获得一个随机字符串
func GetRandomStr(chars []rune, size int) string {
	_len := len(chars)
	result := ""
	for i := 0; i < size; i++ {
		result += string(chars[go_random.RandomInt(_len)])
	}
	return result
}

// GetRandomStr2 获得一个随机字符串，包括大小写、数字
func GetRandomStr2(size int) string {
	joinedString := strings.Join([]string{CapitalLetter, LowercaseLetter, Number}, "")
	return GetRandomStr([]rune(joinedString), size)
}

// GetRandomStr3 获得一个随机字符串，包括大小写、数字、符号
func GetRandomStr3(size int) string {
	joinedString := strings.Join([]string{CapitalLetter, LowercaseLetter, Number, Symbol}, "")
	return GetRandomStr([]rune(joinedString), size)
}

const CapitalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LowercaseLetter = "abcdefghijklmnopqrstuvwxyz"
const Number = "0123456789"
const Symbol = "!@#$%^&*()_+-=;,.:<>?"

// reverseString 字符串翻转
func ReverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Char2Ascii(r rune) int {
	return int(r)
}

// ascii的范围是 [0, 127]
func Ascii2Char(ascii int) string {
	return string(rune(ascii))
}

// 获得大写字母表
func GetUppercaseAlphabet() string {
	result := ""
	for i := 0; i < 26; i++ {
		result = result + string(rune('A'+i))
	}
	return result
}

// 获得小写字母表
func GetLowercaseAlphabet() string {
	result := ""
	for i := 0; i < 26; i++ {
		result = result + string(rune('a'+i))
	}
	return result
}

// RemovePrefix 移除指定前缀
//
//	@param path
//	@param prefix
//	@return string
func RemovePrefix(path, prefix string) string {
	return strings.TrimPrefix(path, prefix)
}

func GetWd() (string, error) {
	return os.Getwd()
}

func SafePasswordEquals(password string, p string) bool {
	if len(password) != len(p) {
		time.Sleep(time.Duration(int(go_random.RandomFloat64() * 100)))
		return false
	}
	var corrent = []rune(password)
	var unknow = []rune(p)
	var flag int32 = 0
	for i := 0; i < len(corrent); i++ {
		flag = flag | (int32(corrent[i]) ^ int32(unknow[i]))
	}
	return flag == 0
}

// 对大小写不敏感的比较两个字符串是否相同
func EqualInsensitivity(a string, b string) bool {
	return strings.EqualFold(a, b)
}
