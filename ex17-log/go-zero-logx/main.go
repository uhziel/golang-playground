package main

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/color"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	ctx := logx.ContextWithFields(context.Background(), logx.Field("bar", 0))

	var cfg logx.LogConf
	conf.FillDefault(&cfg) // LogConf 会被填充上 go-zero 自己 json tag 的默认值
	cfg.Encoding = "plain"
	logx.MustSetup(cfg)
	defer logx.Close()

	logx.AddGlobalFields(logx.Field("foo", 0))
	logx.SetLevel(logx.DebugLevel)

	logx.Debug("hello", 1)
	logx.WithContext(ctx).Debug("hello", 2, logx.WithColor("red", color.FgRed))
	logx.WithDuration(time.Second).Debug("hello", 3)

	logx.ErrorStackf("ErrorStackf:hello %d", 4)
	logx.Severef("Severef:hello %d", 5)
	logx.Statf("Statf:hello %d", 6)
}
