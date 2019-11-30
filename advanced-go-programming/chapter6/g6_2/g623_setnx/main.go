package main

import (
	"github.com/go-redis/redis"
)

//可以使用redis做抢占
func incr() {
	_ = redis.NewClient(&redis.Options{})
}
