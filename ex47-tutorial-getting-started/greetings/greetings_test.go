/*
https://golang.google.cn/doc/tutorial/add-a-test
添加单元测试，会执行 *_test.go 文件中所有 Test* 函数
$ go test
PASS
ok      example.com/greetings   0.001s
$ go test -v
=== RUN   TestHelloName
--- PASS: TestHelloName (0.00s)
=== RUN   TestHelloEmpty
--- PASS: TestHelloEmpty (0.00s)
PASS
ok      example.com/greetings   0.001s
*/

package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "zhulei"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("zhulei")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("zhulei") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
