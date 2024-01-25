package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	input := "2024-01-25T09:15:05.986285209Z [init] Resolving type given VANILLA"
	tStr, logText, found := strings.Cut(input, " ")
	if !found {
		log.Fatalln("log is invalid")
	} else {
		fmt.Printf("Cut(%q, \" \") = %q, %q, %v\n", input, tStr, logText, found)
	}

	t, err := time.Parse(time.RFC3339Nano, tStr)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s'Unix      is %d\n", tStr, t.Unix())
	fmt.Printf("%s'UnixMilli is %d\n", tStr, t.UnixMilli())
	fmt.Printf("%s'UnixMicro is %d\n", tStr, t.UnixMicro())
	fmt.Printf("%s'UnixNano  is %d\n", tStr, t.UnixNano())
}
