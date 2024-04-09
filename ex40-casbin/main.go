package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	enforcer, err := casbin.NewEnforcer("model-acl.conf", "policy.csv")
	if err != nil {
		panic(err)
	}

	result, err := enforcer.Enforce("alice", "data1", "read")
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
