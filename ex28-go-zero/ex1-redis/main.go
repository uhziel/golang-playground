package main

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

func main() {
	conf := redis.RedisConf{
		Host:        "192.168.31.2:6379",
		Type:        "node",
		Pass:        "",
		Tls:         false,
		NonBlock:    false,
		PingTimeout: 3 * time.Second,
	}

	rds := redis.MustNewRedis(conf)
	ctx := context.Background()

	if err := rds.SetCtx(ctx, "foo", "123"); err != nil {
		panic(err)
	}

	v, err := rds.GetCtx(ctx, "foo")
	if err != nil {
		panic(err)
	}

	fmt.Println(v)
}
