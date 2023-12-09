package main

import (
	"context"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	ctx := logx.ContextWithFields(context.Background(), logx.Field("bar", 0))

	var cfg logx.LogConf
	conf.FillDefault(&cfg) // LogConf 会被填充上 go-zero 自己 json tag 的默认值
	logx.MustSetup(cfg)
	defer logx.Close()

	logx.AddGlobalFields(logx.Field("foo", 0))
	logx.SetLevel(logx.DebugLevel)

	logx.Debug("hello", 1)
	logx.WithContext(ctx).Debug("hello", 2)
}
