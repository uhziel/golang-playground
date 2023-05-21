package ex9map

import (
	"testing"
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
