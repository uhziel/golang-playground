package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	nonewline = flag.Bool("n", false, "do not append a newline")
	sep       = flag.String("sep", " ", "separator")
)

var output io.Writer = os.Stdout

func main() {
	flag.Parse()

	if err := echo(!(*nonewline), *sep, flag.Args()); err != nil {
		log.Fatalln(err)
	}
}

func echo(newline bool, sep string, args []string) error {
	var builder strings.Builder
	for i, arg := range args {
		if i > 0 {
			builder.WriteString(sep)
		}
		builder.WriteString(arg)
	}
	if newline {
		builder.WriteByte('\n')
	}
	fmt.Fprint(output, builder.String())

	return nil
}
