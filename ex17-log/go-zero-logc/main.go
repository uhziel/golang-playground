package main

import (
	"context"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	ctx := logx.ContextWithFields(context.Background(), logx.Field("bar", 0))

	cfg := logc.LogConf{}
	logc.MustSetup(cfg)
	defer logc.Close()

	logc.AddGlobalFields(logc.Field("foo", 0))
	logc.SetLevel(logx.ErrorLevel)

	logc.Debug(ctx, "hello", 1)
	logc.Infof(ctx, "hello %d", 2)
	logc.Errorv(ctx, "hello3")
	logc.Sloww(ctx, "hello", logc.Field("cnt", 4))
}
