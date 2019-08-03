package main

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func initPoll(address string, MaxIdle, MaxActive int, IdleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     MaxIdle,
		MaxActive:   MaxActive,
		IdleTimeout: IdleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
