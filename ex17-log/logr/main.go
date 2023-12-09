package main

import (
	"fmt"
	stdlog "log"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
)

func main() {
	//log := stdr.New(stdlog.New(os.Stderr, "", stdlog.LstdFlags|stdlog.Lshortfile))
	logger := stdr.NewWithOptions(
		stdlog.New(os.Stderr, "", stdlog.LstdFlags|stdlog.Lshortfile),
		stdr.Options{
			LogCaller: stdr.All,
		},
	)
	stdr.SetVerbosity(1)
	logger.Info("hello")
	submodule1(logger)
}

func submodule1(logger logr.Logger) {
	sublogger := logger.WithName("submodule1").WithValues("key1", 1)

	sublogger.Info("foobar")
	sublogger.Error(fmt.Errorf("error1"), "foobar2", "key2", 2)
	sublogger.V(1).Info("v1 foobar")
}
