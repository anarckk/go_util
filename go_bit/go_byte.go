package go_bit

import (
	"encoding/binary"
	"log"
)

// Int64ArrayToByteArray 把int64转型成bytes中
//
//	@param i
//	@param _bytes
func Int64ArrayToByteArray(int64Array []int64) []byte {
	size := len(int64Array)
	var _bytes = make([]byte, 0, size*8)
	for i := 0; i < size; i++ {
		i64Bytes := Int64ToBytes(int64Array[i])
		_bytes = append(_bytes, i64Bytes...)
	}
	return _bytes
}

func ByteArrayToInt64Array(_bytes []byte) []int64 {
	iLen := len(_bytes) / 8
	var iArray = make([]int64, iLen)
	for i := 0; i < iLen; i++ {
		log.Println(_bytes[i*8 : 8])
		iArray[i] = BytesToInt64(_bytes[i*8 : i*8+8])
	}
	return iArray
}

func PutUint16(b []byte, v uint16) {
	binary.BigEndian.PutUint16(b, v)
}
func PutUint32(b []byte, v uint32) {
	binary.BigEndian.PutUint32(b, v)
}
func PutUint64(b []byte, v uint64) {
	binary.BigEndian.PutUint64(b, v)
}
func GetUint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}
func GetUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}
func GetUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

type ByteStack []byte

func (b ByteStack) AddInt8(a int8) ByteStack {
	return b
}
func (b ByteStack) AddInt16(a int16) ByteStack {
	return b
}
func (b ByteStack) AddInt32(a int32) ByteStack {
	return b
}
func (b ByteStack) AddInt64(a int64) ByteStack {
	return b
}
func (b ByteStack) AddString(a int8) ByteStack {
	return b
}
