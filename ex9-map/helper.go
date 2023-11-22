package ex9map

import (
	"math/rand"
	"testing"
	"time"
)

type OPType string

const (
	OPSearch OPType = "Search"
	OPAdd    OPType = "Add"
	OPRemove OPType = "Remove"
)

type Check struct {
	Want bool
}

type Call struct {
	OP          OPType
	Arg         int
	Check       *Check
	PrintAfter  bool
	PrintBefore bool
}

func TestMap(t *testing.T, m Map, calls []Call) {
	for i, call := range calls {
		if call.PrintBefore {
			t.Logf("%v", m)
		}
		if call.OP == OPSearch {
			if ans := m.Search(call.Arg); call.Check != nil && call.Check.Want != ans {
				t.Errorf("calls[%d]: m.Search(%d), want %v", i, call.Arg, call.Check.Want)
			}
		} else if call.OP == OPAdd {
			if ans := m.Add(call.Arg); call.Check != nil && call.Check.Want != ans {
				t.Errorf("calls[%d]: m.Add(%d), want %v", i, call.Arg, call.Check.Want)
			}
		} else if call.OP == OPRemove {
			if ans := m.Remove(call.Arg); call.Check != nil && call.Check.Want != ans {
				t.Errorf("calls[%d]: m.Remove(%d), want %v", i, call.Arg, call.Check.Want)
			}
		}
		if call.PrintAfter {
			t.Logf("%v", m)
		}
	}
}

func TestSearch1(t *testing.T, m Map) {
	calls := []Call{
		{OP: OPSearch, Arg: 0, Check: &Check{Want: false}},
		{OP: OPAdd, Arg: 5},
		{OP: OPSearch, Arg: 5, Check: &Check{Want: true}},
		{OP: OPSearch, Arg: 0, Check: &Check{Want: false}},
		{OP: OPAdd, Arg: 4},
		{OP: OPAdd, Arg: 6},
		{OP: OPSearch, Arg: 4, Check: &Check{Want: true}},
		{OP: OPSearch, Arg: 6, Check: &Check{Want: true}},
		{OP: OPSearch, Arg: 7, Check: &Check{Want: false}},
	}

	TestMap(t, m, calls)
}

func TestAdd1(t *testing.T, m Map) {
	calls := []Call{
		{OP: OPAdd, Arg: 5, Check: &Check{Want: true}},
		{OP: OPAdd, Arg: 5, Check: &Check{Want: false}},
		{OP: OPAdd, Arg: 3, Check: &Check{Want: true}},
		{OP: OPAdd, Arg: 3, Check: &Check{Want: false}},
		{OP: OPAdd, Arg: 9, Check: &Check{Want: true}},
		{OP: OPAdd, Arg: 9, Check: &Check{Want: false}},
		{OP: OPAdd, Arg: 1, Check: &Check{Want: true}},
		{OP: OPAdd, Arg: 1, Check: &Check{Want: false}},
		{OP: OPAdd, Arg: 8, Check: &Check{Want: true}},
		{OP: OPAdd, Arg: 8, Check: &Check{Want: false}},
		{OP: OPAdd, Arg: 6, Check: &Check{Want: true}},
		{OP: OPAdd, Arg: 6, Check: &Check{Want: false}},
		{OP: OPAdd, Arg: 4, Check: &Check{Want: true}},
		{OP: OPAdd, Arg: 4, Check: &Check{Want: false}},
		{OP: OPSearch, Arg: 5, Check: &Check{Want: true}, PrintBefore: true},
		{OP: OPSearch, Arg: 2, Check: &Check{Want: false}},
		{OP: OPSearch, Arg: 6, Check: &Check{Want: true}},
		{OP: OPSearch, Arg: 7, Check: &Check{Want: false}},
		{OP: OPSearch, Arg: 1, Check: &Check{Want: true}},
	}

	TestMap(t, m, calls)
}

func TestRemove1(t *testing.T, m Map) {
	calls := []Call{
		{OP: OPAdd, Arg: 1},
		{OP: OPAdd, Arg: 2},
		{OP: OPAdd, Arg: 3},
		{OP: OPSearch, Arg: 0, Check: &Check{Want: false}},
		{OP: OPAdd, Arg: 4, PrintAfter: true},
		{OP: OPSearch, Arg: 1, Check: &Check{Want: true}},
		{OP: OPRemove, Arg: 0, Check: &Check{Want: false}},
		{OP: OPRemove, Arg: 1, Check: &Check{Want: true}},
		{OP: OPSearch, Arg: 1, Check: &Check{Want: false}},
		{OP: OPRemove, Arg: 3, Check: &Check{Want: true}},
		{OP: OPRemove, Arg: 4, Check: &Check{Want: true}},
		{OP: OPRemove, Arg: 2, Check: &Check{Want: true}, PrintAfter: true},
		{OP: OPRemove, Arg: 2, Check: &Check{Want: false}},
	}
	TestMap(t, m, calls)
}

func TestRemove2(t *testing.T, m Map) {
	// curNode.Left != nil && curNode.Right == nil
	// parentNode != nil && parentNode.Right == curNode
	calls := []Call{
		{OP: OPAdd, Arg: 0},
		{OP: OPAdd, Arg: 5},
		{OP: OPAdd, Arg: 2},
		{OP: OPAdd, Arg: 1},
		{OP: OPSearch, Arg: 0, Check: &Check{Want: true}},
		{OP: OPRemove, Arg: 5, PrintBefore: true, PrintAfter: true},
		{OP: OPSearch, Arg: 2, Check: &Check{Want: true}},
		{OP: OPSearch, Arg: 3, Check: &Check{Want: false}},
	}
	TestMap(t, m, calls)
}

func TestRemove3(t *testing.T, m Map) {
	// curNode.Left != nil && curNode.Right == nil
	// parentNode != nil && parentNode.Left == curNode
	calls := []Call{
		{OP: OPAdd, Arg: 6},
		{OP: OPAdd, Arg: 4},
		{OP: OPAdd, Arg: 9},
		{OP: OPAdd, Arg: 3},
		{OP: OPRemove, Arg: 4, PrintBefore: true, PrintAfter: true},
		{OP: OPSearch, Arg: 3, Check: &Check{Want: true}},
	}
	TestMap(t, m, calls)
}

func TestRemove4(t *testing.T, m Map) {
	// curNode.Left == nil && curNode.Right != nil
	// parentNode != nil && parentNode.Left == curNode
	calls := []Call{
		{OP: OPAdd, Arg: 6},
		{OP: OPAdd, Arg: 4},
		{OP: OPAdd, Arg: 9},
		{OP: OPAdd, Arg: 5},
		{OP: OPRemove, Arg: 4, PrintBefore: true, PrintAfter: true},
		{OP: OPSearch, Arg: 5, Check: &Check{Want: true}},
		{OP: OPSearch, Arg: 4, Check: &Check{Want: false}},
	}
	TestMap(t, m, calls)
}

func TestRemove5(t *testing.T, m Map) {
	// curNode.Left != nil && curNode.Right != nil
	// parentNode != nil && parentNode.Left == curNode
	calls := []Call{
		{OP: OPAdd, Arg: 6},
		{OP: OPAdd, Arg: 4},
		{OP: OPAdd, Arg: 9},
		{OP: OPAdd, Arg: 3},
		{OP: OPAdd, Arg: 5},
		{OP: OPRemove, Arg: 4, PrintBefore: true, PrintAfter: true},
		{OP: OPSearch, Arg: 5, Check: &Check{Want: true}},
		{OP: OPSearch, Arg: 4, Check: &Check{Want: false}},
	}
	TestMap(t, m, calls)
}

func TestRemove6(t *testing.T, m Map) {
	// curNode.Left != nil && curNode.Right != nil
	// parentNode == nil
	calls := []Call{
		{OP: OPAdd, Arg: 8},
		{OP: OPAdd, Arg: 9},
		{OP: OPAdd, Arg: 1},
		{OP: OPAdd, Arg: 7},
		{OP: OPAdd, Arg: 6},
		{OP: OPAdd, Arg: 5},
		{OP: OPRemove, Arg: 8, PrintBefore: true, PrintAfter: true},
		{OP: OPSearch, Arg: 5, Check: &Check{Want: true}},
		{OP: OPSearch, Arg: 4, Check: &Check{Want: false}},
	}
	TestMap(t, m, calls)
}

func TestRemove7(t *testing.T, m Map) {
	// curNode.Left != nil && curNode.Right != nil
	// parentNode != nil && parentNode.Right == curNode
	calls := []Call{
		{OP: OPAdd, Arg: 1},
		{OP: OPAdd, Arg: 8},
		{OP: OPAdd, Arg: 9},
		{OP: OPAdd, Arg: 2},
		{OP: OPAdd, Arg: 7},
		{OP: OPAdd, Arg: 6},
		{OP: OPAdd, Arg: 5},
		{OP: OPRemove, Arg: 8, PrintBefore: true, PrintAfter: true},
		{OP: OPSearch, Arg: 5, Check: &Check{Want: true}},
		{OP: OPSearch, Arg: 4, Check: &Check{Want: false}},
		{OP: OPSearch, Arg: 2, Check: &Check{Want: true}},
	}
	TestMap(t, m, calls)
}

func TestRandom1(t *testing.T, m Map) {
	seed := time.Now().UnixNano()
	t.Logf("seed: %d", seed)
	rnd := rand.New(rand.NewSource(seed))
	n := 10
	nodes := rnd.Perm(n)
	calls := []Call{}
	for _, node := range nodes {
		calls = append(calls, Call{OP: OPAdd, Arg: node})
	}

	calls = append(calls, []Call{
		{OP: OPSearch, Arg: nodes[0], Check: &Check{Want: true}, PrintBefore: true},
		{OP: OPSearch, Arg: nodes[(len(nodes)-1)/2], Check: &Check{Want: true}},
		{OP: OPSearch, Arg: nodes[len(nodes)-1], Check: &Check{Want: true}},
		{OP: OPSearch, Arg: n + 100, Check: &Check{Want: false}},
		{OP: OPRemove, Arg: nodes[(len(nodes)-1)/2], Check: &Check{Want: true}},
		{OP: OPSearch, Arg: nodes[(len(nodes)-1)/2], Check: &Check{Want: false}},
	}...)

	TestMap(t, m, calls)
}
