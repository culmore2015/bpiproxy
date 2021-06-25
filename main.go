package main

import (
	"net/http"
	"proxyblockpi/bpi"
	"proxyblockpi/db/redis"
)
var GRedisPool *redis.ConnPool

func main() {
	GRedisPool = redis.InitRedisPool("127.0.0.1:9944","", 0, 20, 15)
	bpi.LoadServices()
	handler  := &handler{}
	_ = http.ListenAndServe(":9999", handler)
}