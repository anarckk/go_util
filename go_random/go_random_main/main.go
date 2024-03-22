package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"gitea.bee.anarckk.me/anarckk/go_util/go_bit"
	"gitea.bee.anarckk.me/anarckk/go_util/go_func"
	"gitea.bee.anarckk.me/anarckk/go_util/go_random"
	"gitea.bee.anarckk.me/anarckk/go_util/go_unit"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
func test2() {
	log.Println(go_func.Combine2(go_random.RandomBytes1, go_bit.BytesToBytesStr)(15))
	log.Println(go_func.Combine2(go_random.RandomBytes2, go_bit.BytesToBytesStr)(15))
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func TestRandomBool() {
	go_random.InitSeed()
	var total = 10_0000_0000
	var result = make([]bool, 0)
	for i := 0; i < total; i++ {
		result = append(result, go_random.RandomBool(0.1))
	}
	var count int = 0
	for _, _r := range result {
		if _r {
			count++
		}
	}
	fmt.Println(float64(count) / float64(total))
}

func Test3() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		var total int64 = 0
		for {
			if ctx.Err() != nil {
				log.Println(total)
				log.Println(go_unit.HumanReadableByteCountBin(total))
				break
			}
			_bytes := go_random.RandomBytes2(1024)
			total += int64(len(_bytes))
		}
	}()
	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)
}

func main() {
	Test3()
}
