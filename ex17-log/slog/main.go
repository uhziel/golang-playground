package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()
	log.Println("0hello world from log package")
	slog.Info("1hello world")
	slog.Info("2hello world", "tag", "golang")
	slog.LogAttrs(ctx, slog.LevelInfo, "3hello world", slog.String("tag", "golang"))

	// Default Logger
	logger := slog.Default()
	logger.Info("4hello world", "tag", "golang")

	// Text Logger logfmt
	textlogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		//AddSource: true,
		Level: slog.LevelDebug,
	}))
	textlogger.Debug("5hello world debug", "tag", "golang")
	textlogger.Info("5hello world info", slog.Group("request", "method", "GET", "url", "http://example.com"))

	// JSON Logger
	jsonlogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonlogger.Info("6hello world", "tag", "golang")

	// Logger.With
	submodule(logger.With("module", "module1"))
	submodule(logger.With(slog.String("module", "module2")))

	// Logger.WithGroup
	submodule(logger.WithGroup("fight"))

	// LogValuer
	slog.Info("register", "user", UserInfo{"user1", "123456"})
	slog.Info("register", "userinfogroup", UserInfoGroup{"user1", "123456"})

	// Group
	jsonlogger.Info("7hello world", slog.Group("request", "method", "GET", "url", "http://example.com"))
}

type UserInfo struct {
	Username, Password string
}

func (u UserInfo) LogValue() slog.Value {
	return slog.StringValue(fmt.Sprintf("{user:%s,password:*}", u.Username))
}

type UserInfoGroup struct {
	Username, Password string
}

func (u UserInfoGroup) LogValue() slog.Value {
	return slog.GroupValue(slog.String("user", u.Username), slog.String("password", "*"))
}

func submodule(log *slog.Logger) {
	log.Info("7hello world", "tag", "golang")
}
