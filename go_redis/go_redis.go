/*
 * @Author: anarckk anarckk@gmail.com
 * @Date: 2023-12-16 22:44:32
 * @LastEditTime: 2023-12-18 08:56:02
 * @Description: go_redis工具
 *
 * Copyright (c) 2023 by anarckk, All Rights Reserved.
 */
package go_redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Connection redis.Client

// NewConnection 新建连接
// 需要注意的是，需要手动关闭
func NewConnection(address string) *Connection {
	return (*Connection)(redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	}))
}

func (conn *Connection) client() *redis.Client {
	return (*redis.Client)(conn)
}

func (conn *Connection) Close() error {
	return (*redis.Client)(conn).Close()
}

func (conn *Connection) Set(ctx context.Context, key string, value string) error {
	return conn.client().Set(ctx, key, value, 0).Err()
}

func (conn *Connection) Get(ctx context.Context, key string) (string, error) {
	return conn.client().Get(ctx, key).Result()
}

func (conn *Connection) Del(ctx context.Context, key string) error {
	_, err := conn.client().Del(ctx, key).Result()
	return err
}

func (conn *Connection) HSet(ctx context.Context, key string, field string, value string) error {
	_, err := conn.client().HSet(ctx, key, field, value).Result()
	return err
}

func (conn *Connection) HGet(ctx context.Context, key string, field string) (string, error) {
	return conn.client().HGet(ctx, key, field).Result()
}

func (conn *Connection) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return conn.client().Expire(ctx, key, expiration).Result()
}

// TTL 返回key的剩余时间
func (conn *Connection) TTL(ctx context.Context, key string) (time.Duration, error) {
	return conn.client().TTL(ctx, key).Result()
}

type Sub struct {
	pubsub *redis.PubSub
}

func (sub *Sub) Channal() <-chan *redis.Message {
	return sub.pubsub.Channel()
}

func (sub *Sub) Close() error {
	return sub.pubsub.Close()
}

func (conn *Connection) Subscribe(ctx context.Context, channel string) *Sub {
	pubsub := conn.client().Subscribe(ctx, channel)
	return &Sub{pubsub}
}

func (conn *Connection) Publish(ctx context.Context, channel string, msg string) error {
	return conn.client().Publish(ctx, channel, msg).Err()
}
