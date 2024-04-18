package main

import (
	"context"
	"fmt"
	"log"

	"github.com/uhziel/golang-playground/ex41-context-withvalue/user"
)

func main() {
	u := user.User{
		ID:   "10234",
		Name: "name",
	}

	ctx := user.NewContext(context.Background(), &u)
	print(ctx)
}

func print(ctx context.Context) {
	u, ok := user.FromContext(ctx)
	if !ok {
		log.Fatalln("not found")
	}
	fmt.Printf("%#v\n", u)
}
