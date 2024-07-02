/*
//https://golang.org/doc/tutorial/getting-started
# 一、初始化 module
$ go mod init example.com/hello
go: creating new go.mod: module example.com/hello
$ cat go.mod # go.mod 的内容非常简单，只有 module 的名字（也是其下载地址）还有使用的 go 版本
module example.com/hello

go 1.16

# 二、编写 package main 并且执行
$ go run . # package main 就是本文件提供的
Hello, World!

# 三、使用外部的包
import	"rsc.io/quote" 后
$ go run . # 改了 hello.go 添加上面的 import 语句后，发现没有并没有该包，run 命令不会自动下载该包
hello.go:5:8: no required module provides package rsc.io/quote; to add it:
        go get rsc.io/quote
$ go mod tidy # 这个才会自动下载该包，生成 go.sum 文件(用来后期做 checksum)、更新 go.mod
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
$ cat go.sum
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZOaTkIIMiVjBQcw93ERBE4m30iBm00nkL0i8=
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPXsUe+TKr0=
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9JXDnKaTXpA=
$ cat go.mod # 有更新 require 指明需要 rsc.io/quote。 go.mod、go.sum 的作用有点类似 npm 的 package.json、package-lock.json。
module example.com/hello

go 1.16

require rsc.io/quote v1.5.2
*/

/*
#https://golang.google.cn/doc/tutorial/call-module-code
# 使用本地编写的包
# 一、代码中导入依赖关系
# import "example.com/greetings"
# 二、hello module 编辑replace
$ go mod edit -replace=example.com/greetings=../greetings
$ go mod tidy
module example.com/hello

go 1.13

require (
        example.com/greetings v0.0.0-00010101000000-000000000000
        rsc.io/quote v1.5.2
)

replace example.com/greetings => ../greetings
*/

package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println(quote.Go())
	names := []string{
		"zhulei",
		"lilei",
	}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal("your name is empty")
	}
	fmt.Println(messages)
}
