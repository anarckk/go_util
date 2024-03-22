/*
 * @Author: anarckk anarckk@gmail.com
 * @Date: 2023-09-03 21:58:42
 * @LastEditTime: 2024-01-03 10:38:21
 * @Description: bit操作工具类
 *
 * Copyright (c) 2023 by anarckk, All Rights Reserved.
 */
package go_bit

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func ByteToByteStr(b byte) string {
	return ByteToByteStr2(b)
}

func ByteToByteStr1(b byte) string {
	var str = ""
	for i := 7; i >= 0; i-- {
		str = str + strconv.Itoa(int((b>>i)&0x01))
	}
	return str
}

func ByteToByteStr2(b byte) string {
	var str = ""
	str += strconv.Itoa(int((b >> 7) & 0x01))
	str += strconv.Itoa(int((b >> 6) & 0x01))
	str += strconv.Itoa(int((b >> 5) & 0x01))
	str += strconv.Itoa(int((b >> 4) & 0x01))
	str += strconv.Itoa(int((b >> 3) & 0x01))
	str += strconv.Itoa(int((b >> 2) & 0x01))
	str += strconv.Itoa(int((b >> 1) & 0x01))
	str += strconv.Itoa(int((b) & 0x01))
	return str
}

func ByteStrToByte(binaryString string) byte {
	var result byte
	for i := 0; i < 8; i++ {
		c := binaryString[i]
		if c == '1' {
			result |= (1 << (7 - i))
		}
	}
	return result
}

func BytesToBytesStr(bytes []byte) string {
	var str = ""
	for _, b := range bytes {
		str = str + ByteToByteStr(b) + " "
	}
	return str
}

func BytesStrToBytes(binaryString string) []byte {
	var strArr = strings.Split(binaryString, " ")
	var bytes = make([]byte, 0, len(strArr))
	for _, str := range strArr {
		var b byte = ByteStrToByte(str)
		bytes = append(bytes, b)
	}
	return bytes
}

func ByteToHexStr(b byte) string {
	return ByteToHexStr2(b)
}

func ByteToHexStr1(b byte) string {
	return fmt.Sprintf("%02X", b)
}

func ByteToHexStr2(b byte) string {
	hexDigits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}
	result := make([]rune, 2)
	value := int(b) & 0xFF
	result[0] = hexDigits[value>>4]
	result[1] = hexDigits[value&0x0F]
	return string(result)
}

func HexStrToByte(hexStr string) byte {
	high := int(hexStr[0]) - '0'
	if high > 9 {
		high = int(hexStr[0]) - 'A' + 10
	}
	low := int(hexStr[1]) - '0'
	if low > 9 {
		low = int(hexStr[1]) - 'A' + 10
	}
	return byte((high << 4) + low)
}

// 比较紧凑的格式 00038D7EA4C67FFF
func BytesToHexStrCompact(_bytes []byte) string {
	return strings.ToUpper(hex.EncodeToString(_bytes))
}

func HexStrCompactToBytes(hexCompactStr string) []byte {
	_bytes, err := hex.DecodeString(strings.ToLower(hexCompactStr))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return _bytes
}

// 比较宽松的格式 00 03 8D 7E A4 C6 7F FF
func BytesToHexStr(bytes []byte) string {
	var str = ""
	for _, b := range bytes {
		str = str + ByteToHexStr(b) + " "
	}
	return strings.Trim(str, " ")
}

func HexStrToBytes(hexStr string) []byte {
	subHex := strings.Split(hexStr, " ")
	var bytes = make([]byte, 0, len(subHex))
	for _, _sub := range subHex {
		bytes = append(bytes, HexStrToByte(_sub))
	}
	return bytes
}

func BytesStrToHexStr(bytesStr string) string {
	return BytesToHexStr(BytesStrToBytes(bytesStr))
}

func HexStrToBytesStr(hexStr string) string {
	return BytesToBytesStr(HexStrToBytes(hexStr))
}

func CheckIsAsciiNumber(c rune) bool {
	return c >= 48 && c <= 57
}

func NumberToascii(b int) rune {
	return rune(b + 48)
}

func AsciiToNumber(c rune) int {
	return int(c - 48)
}

func Int16ToBytes(a int16) []byte {
	bytes := make([]byte, 2)
	bytes[0] = byte(a >> 8)
	bytes[1] = byte(a)
	return bytes
}

func BytesToInt16(bytes []byte) int16 {
	var a int16 = 0
	a |= int16(bytes[0]&0xff) << 8
	a |= int16(bytes[1] & 0xff)
	return a
}

func Uint16ToBytes(a uint16) []byte {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, a)
	return bytes
}

func BytesToUint16(bytes []byte) uint16 {
	return binary.BigEndian.Uint16(bytes[0:2])
}

func IntToBytes(a int) []byte {
	bytes := make([]byte, 4)
	bytes[0] = byte(a >> 24)
	bytes[1] = byte(a >> 16)
	bytes[2] = byte(a >> 8)
	bytes[3] = byte(a)
	return bytes
}

func BytesToInt(bytes []byte) int {
	size := 32
	step := size / 8
	a := 0
	for i := 0; i < step; i++ {
		flag := int(bytes[i]&0xff) << ((step - 1 - i) * 8)
		a |= flag
	}
	return a
}

func Uint32ToBytes(a uint32) []byte {
	return Uint32ToBytesTwo(a)
}

func Uint32ToBytesOne(a uint32) []byte {
	bytes := make([]byte, 4)
	bytes[0] = byte(a >> 24)
	bytes[1] = byte(a >> 16)
	bytes[2] = byte(a >> 8)
	bytes[3] = byte(a)
	return bytes
}

func Uint32ToBytesTwo(a uint32) []byte {
	_bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(_bytes[0:4], a)
	return _bytes
}

func BytesToUint32(bytes []byte) uint32 {
	return BytesToUint32Two(bytes)
}

func BytesToUint32One(bytes []byte) uint32 {
	size := 32
	step := size / 8
	var a uint32 = 0
	for i := 0; i < step; i++ {
		flag := uint32(bytes[i]&0xff) << ((step - 1 - i) * 8)
		a |= flag
	}
	return a
}

func BytesToUint32Two(bytes []byte) uint32 {
	return binary.BigEndian.Uint32(bytes[:4])
}

func Int64ToBytes(a int64) []byte {
	size := 64
	step := size / 8
	bytes := make([]byte, step)
	for i := 0; i < step; i++ {
		bytes[i] = byte(a >> ((step - 1 - i) * 8))
	}
	return bytes
}

func BytesToInt64(bytes []byte) int64 {
	size := 64
	step := size / 8
	var a int64 = 0
	for i := 0; i < step; i++ {
		flag := int64(bytes[i]&0xff) << ((step - 1 - i) * 8)
		a |= flag
	}
	return a
}

func Uint64ToBytes(a uint64) []byte {
	return Uint64ToBytes2(a)
}
func Uint64ToBytes1(a uint64) []byte {
	size := 64
	step := size / 8
	bytes := make([]byte, step)
	for i := 0; i < step; i++ {
		bytes[i] = byte(a >> ((step - 1 - i) * 8))
	}
	return bytes
}
func Uint64ToBytes2(a uint64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, a)
	return bytes
}

func BytesToUint64(bytes []byte) uint64 {
	return BytesToUint64_2(bytes)
}
func BytesToUint64_1(bytes []byte) uint64 {
	size := 64
	step := size / 8
	var a uint64 = 0
	for i := 0; i < step; i++ {
		flag := uint64(bytes[i]&0xff) << ((step - 1 - i) * 8)
		a |= flag
	}
	return a
}
func BytesToUint64_2(bytes []byte) uint64 {
	return binary.BigEndian.Uint64(bytes)
}

func RuneToBytes(c rune) []byte {
	return RuneToBytes2(c)
}

// Deprecated: 这个做法不对,获得到的数据不是utf8的byte[]
func RuneToBytes1(c rune) []byte {
	return Uint32ToBytes(uint32(c))
}

func RuneToBytes2(c rune) []byte {
	return []byte(string(c))
}

func BytesToRune(bytes []byte) rune {
	return BytesToRune2(bytes)
}

// Deprecated: 这个做法不对,获得到的数据不是utf8的byte[]
func BytesToRune1(bytes []byte) rune {
	return rune(BytesToUint32(bytes))
}

func BytesToRune2(bytes []byte) rune {
	r := []rune(string(bytes))
	return r[0]
}
