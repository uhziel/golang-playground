package chaininghashtable

import (
	"testing"

	ex9map "github.com/uhziel/golang-playground/ex9-map"
)

func TestSearch1(t *testing.T) {
	m := New()
	ex9map.TestSearch1(t, m)
}

func TestAdd1(t *testing.T) {
	m := New()
	ex9map.TestAdd1(t, m)
}

func TestRemove1(t *testing.T) {
	m := New()
	ex9map.TestRemove1(t, m)
}

func TestRemove2(t *testing.T) {
	m := New()
	ex9map.TestRemove2(t, m)
}

func TestRemove3(t *testing.T) {
	m := New()
	ex9map.TestRemove3(t, m)
}

func TestRemove4(t *testing.T) {
	m := New()
	ex9map.TestRemove4(t, m)
}

func TestRemove5(t *testing.T) {
	m := New()
	ex9map.TestRemove5(t, m)
}

func TestRemove6(t *testing.T) {
	m := New()
	ex9map.TestRemove6(t, m)
}

func TestRemove7(t *testing.T) {
	m := New()
	ex9map.TestRemove7(t, m)
}

func TestRandom1(t *testing.T) {
	m := New()
	ex9map.TestRandom1(t, m)
}
