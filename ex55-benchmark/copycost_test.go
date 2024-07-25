package copycost

import (
	"testing"
)

func BenchmarkCopyForI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copyForI()
	}
}

func BenchmarkCopyForR1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copyForR1()
	}
}

func BenchmarkCopyForR2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copyForR2()
	}
}
