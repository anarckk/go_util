/**
 * Author: anarckk anarckk@gmail.com
 * Date: 2023-05-04 13:44:58
 * LastEditTime: 2023-05-04 14:12:51
 * Description: 随机数工具
 *
 * Copyright (c) 2023 by anarckk, All Rights Reserved.
 */
package go_random

import (
	cryptoRand "crypto/rand"
	"log"
	"math/rand"
	"time"

	"gitea.bee.anarckk.me/anarckk/go_util/go_bit"
)

// Deprecated: 废弃
func InitSeed() {
	rand.Seed(time.Now().Unix())
}

// RandomInt 生成随机整数，范围 [0, max)
//
//	@param max
//	@return int
func RandomInt(max int) int {
	return rand.Intn(max)
}

func RandomInt64() int64 {
	return rand.Int63()
}

// RandomInt2 生成随机整数，范围 [min, max)
//
//	@param min
//	@param max
//	@return int
func RandomInt2(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomFloat64 生成随机浮点数
//
//	@return float64 [0.0,1.0)
func RandomFloat64() float64 {
	return rand.Float64()
}

// RandomBool 根据给定概率生成随机bool值
//
//	@param probability 为true的概率多大
//	@return bool
func RandomBool(probability float64) bool {
	_random := rand.Float64()
	return _random <= probability
}

func RandomBytes(size int) []byte {
	return RandomBytes2(size)
}

// 一秒钟最快能获得 384720896 366.9 MiB
func RandomBytes1(size int) []byte {
	iv := make([]byte, size)
	_, err := cryptoRand.Read(iv)
	if err != nil {
		log.Println(err)
		return iv
	}
	return iv
}

// 一秒钟最快能获得 988253184 942.5 MiB
func RandomBytes2(size int) []byte {
	var result = make([]byte, 0, size)
	var last = size
	for {
		if last >= 8 {
			ran8Bytes := go_bit.Uint64ToBytes(rand.Uint64())
			result = append(result, ran8Bytes...)
			last -= 8
		} else if last > 0 {
			ran4Bytes := go_bit.Uint32ToBytes(rand.Uint32())
			step := min(last, 4)
			for i := 0; i < step; i++ {
				result = append(result, ran4Bytes[i])
			}
			last -= step
		} else if last <= 0 {
			break
		}
	}
	return result
}
