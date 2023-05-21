package bst

import (
	"testing"

	ex9map "github.com/uhziel/golang-playground/cmd/ex9-map"
)

func TestCase1(t *testing.T) {
	m := New()
	calls := []ex9map.Call{
		{OP: ex9map.OPAdd, Arg: 1},
		{OP: ex9map.OPAdd, Arg: 2},
		{OP: ex9map.OPAdd, Arg: 3},
		{OP: ex9map.OPSearch, Arg: 0, Check: &ex9map.Check{Want: false}},
		{OP: ex9map.OPAdd, Arg: 4, PrintAfter: true},
		{OP: ex9map.OPSearch, Arg: 1, Check: &ex9map.Check{Want: true}},
		{OP: ex9map.OPRemove, Arg: 0},
		{OP: ex9map.OPRemove, Arg: 1},
		{OP: ex9map.OPSearch, Arg: 1, Check: &ex9map.Check{Want: false}},
	}
	ex9map.TestMap(t, m, calls)
}

func TestCase2(t *testing.T) {
	m := New()
	calls := []ex9map.Call{
		{OP: ex9map.OPAdd, Arg: 0},
		{OP: ex9map.OPAdd, Arg: 5},
		{OP: ex9map.OPAdd, Arg: 2},
		{OP: ex9map.OPAdd, Arg: 1},
		{OP: ex9map.OPSearch, Arg: 0, Check: &ex9map.Check{Want: true}},
		{OP: ex9map.OPRemove, Arg: 5, PrintBefore: true, PrintAfter: true},
		{OP: ex9map.OPSearch, Arg: 2, Check: &ex9map.Check{Want: true}},
		{OP: ex9map.OPSearch, Arg: 3, Check: &ex9map.Check{Want: false}},
	}
	ex9map.TestMap(t, m, calls)
}

func TestCase3(t *testing.T) {
	m := New()
	calls := []ex9map.Call{
		{OP: ex9map.OPAdd, Arg: 6},
		{OP: ex9map.OPAdd, Arg: 4},
		{OP: ex9map.OPAdd, Arg: 9},
		{OP: ex9map.OPAdd, Arg: 3},
		{OP: ex9map.OPRemove, Arg: 4, PrintBefore: true, PrintAfter: true},
		{OP: ex9map.OPSearch, Arg: 3, Check: &ex9map.Check{Want: true}},
	}
	ex9map.TestMap(t, m, calls)
}

func TestCase4(t *testing.T) {
	tree := New()
	tree.Add(6)
	tree.Add(4)
	tree.Add(9)
	tree.Add(5)
	tree.Remove(4)
	t.Logf("%v", tree)
}

func TestCase5(t *testing.T) {
	tree := New()
	tree.Add(6)
	tree.Add(4)
	tree.Add(9)
	tree.Add(3)
	tree.Add(5)
	tree.Remove(4)
	t.Logf("%v", tree)
}

func TestCase6(t *testing.T) {
	tree := New()
	tree.Add(8)
	tree.Add(9)
	tree.Add(1)
	tree.Add(7)
	tree.Add(6)
	tree.Add(5)
	t.Logf("%v", tree)
	tree.Remove(8)
	t.Logf("%v", tree)
}

func TestCase7(t *testing.T) {
	tree := New()
	tree.Add(8)
	tree.Add(9)
	tree.Add(7)
	tree.Add(6)
	tree.Add(5)
	t.Logf("%v", tree)
	tree.Remove(8)
	t.Logf("%v", tree)
}

func TestCase8(t *testing.T) {
	tree := New()
	tree.Add(1)
	tree.Add(8)
	tree.Add(9)
	tree.Add(2)
	tree.Add(7)
	tree.Add(6)
	tree.Add(5)
	t.Logf("%v", tree)
	tree.Remove(8)
	t.Logf("%v", tree)
}

func TestCase9(t *testing.T) {
	tree := New()
	tree.Add(1)
	tree.Add(1)
	t.Logf("%v", tree)
	tree.Remove(8)
	t.Logf("%v", tree)
}
