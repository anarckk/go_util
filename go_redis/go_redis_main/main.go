package main

import (
	"context"
	"log"
	"time"

	"gitea.bee.anarckk.me/anarckk/go_util/go_redis"
	"gitea.bee.anarckk.me/anarckk/go_util/go_time"
)

func TestPubsub() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	go func() {
		conn := go_redis.NewConnection("go_redis:6379")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		sub := conn.Subscribe(ctx, "time")
	FF:
		for {
			select {
			case c := <-sub.Channal():
				log.Println(c.Payload)
			case <-ctx.Done():
				sub.Close()
				break FF
			}
		}
	}()
	go func() {
		conn := go_redis.NewConnection("go_redis:6379")
		ctx := context.Background()
		for {
			conn.Publish(ctx, "time", go_time.CurrentDatetimeStr())
			time.Sleep(time.Second)
		}
	}()
	log.Println("两个协程已启动")
	select {}
}

func main() {
	TestPubsub()
}
