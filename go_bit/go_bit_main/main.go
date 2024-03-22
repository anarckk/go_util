package main

import (
	"log"

	"gitea.bee.anarckk.me/anarckk/go_util/go_bit"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
func test1() {
	var a int64 = 777771234123415
	log.Println(go_bit.BytesToBytesStr(go_bit.Int64ToBytes(a)))
	var b int64 = 999888823423423499
	log.Println(go_bit.BytesToBytesStr(go_bit.Int64ToBytes(b)))
	var arr = []int64{a, b}
	_bytes := go_bit.Int64ArrayToByteArray(arr)
	log.Println(go_bit.BytesToBytesStr(_bytes))
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func test2() {
	var a uint64 = 9694123453458919323
	log.Println(go_bit.Uint64ToBytes1(a))
	log.Println(go_bit.Uint64ToBytes2(a))
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func TestByteCountBinary() {
	log.Println(go_bit.ByteToByteStr(255))
	log.Println(go_bit.ByteStrToByte("11111111"))

	var b = []byte{1, 3, 15, 255}
	log.Println(go_bit.BytesToBytesStr(b))
	var c = go_bit.BytesStrToBytes("00000001 00000011 00001111 11111111")
	log.Println(c)

	log.Println(go_bit.ByteToHexStr1(0b11111101))
	log.Println(go_bit.ByteToHexStr2(0b11111101))
	log.Println(go_bit.HexStrToByte("A7"))

	log.Println(go_bit.BytesToHexStr(b))
	log.Println(go_bit.HexStrToBytes("01 03 0F FF"))

	log.Println(go_bit.BytesStrToHexStr("00000001 00000011 00001111 11111111"))
	log.Println(go_bit.HexStrToBytesStr("01 03 0F FF"))

	log.Println(go_bit.BytesToBytesStr(go_bit.RuneToBytes('纞')))
	log.Println(go_bit.BytesToBytesStr(go_bit.RuneToBytes('鸿')))
	log.Println(go_bit.BytesToBytesStr([]byte("鸿"))) // 11101001 10111000 10111111
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// TestByteCountBinary()
	hexStr := go_bit.BytesToHexStrCompact(go_bit.Int64ToBytes(99999999999998))
	log.Println(hexStr)
	i64 := go_bit.BytesToInt64(go_bit.HexStrCompactToBytes(hexStr))
	log.Println(i64)
	var b2 byte = 0b10011011
	log.Println(go_bit.ByteToByteStr2(b2))
}
