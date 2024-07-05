package word

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPalindrome(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"abc", false},
		{"aba", true},
		{"", true},
		{"foobar", false},
		{"refer", true},
		{"été", true},
		{"天连水尾水连天", true},
		{"A man, a plan, a canal: Panama", true},
	}

	for _, c := range cases {
		if res := IsPalindrome(c.input); res != c.want {
			t.Errorf("IsPalindrome(%q) = %v. want:%v", c.input, res, c.want)
		}
	}
}

func randomPalindrome(rng *rand.Rand) string {
	l := rng.Intn(25)
	runes := make([]rune, l)
	for i := 0; i < (l+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[l-i-1] = r
	}
	return string(runes)
}

func TestRandomPalindrome(t *testing.T) {
	seed := time.Now().UnixNano()
	t.Logf("seed:%d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 10; i++ {
		input := randomPalindrome(rng)
		if !IsPalindrome(input) {
			t.Errorf("IsPalindrome(%q) = false", input)
		}
	}
}

func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("hello"))
	// Output:
	//true
	//false
}
