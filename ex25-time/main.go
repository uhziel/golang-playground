package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	tStr := "2024-01-25T02:51:42.691224434Z"

	t, err := time.Parse(time.RFC3339Nano, tStr)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s'Unix      is %d\n", tStr, t.Unix())
	fmt.Printf("%s'UnixMilli is %d\n", tStr, t.UnixMilli())
	fmt.Printf("%s'UnixMicro is %d\n", tStr, t.UnixMicro())
	fmt.Printf("%s'UnixNano  is %d\n", tStr, t.UnixNano())
}
