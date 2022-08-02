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
	log := stdr.NewWithOptions(
		stdlog.New(os.Stderr, "", stdlog.LstdFlags|stdlog.Lshortfile),
		stdr.Options{
			LogCaller: stdr.All,
		},
	)
	stdr.SetVerbosity(1)
	log.Info("hello")
	submodule1(log)
}

func submodule1(log logr.Logger) {
	sublog := log.WithName("submodule1").WithValues("key1", 1)

	sublog.Info("foobar")
	sublog.Error(fmt.Errorf("error1"), "foobar2", "key2", 2)
	sublog.V(1).Info("v1 foobar")
}
