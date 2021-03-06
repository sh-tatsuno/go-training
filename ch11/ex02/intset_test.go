package intset

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

func Test_Add_Has(t *testing.T) {
	i := IntSet{}
	m := newMapIntSet()

	if !checkHas(i, m) {
		t.Errorf("err. mismatch i & m")
	}
	if i.String() != m.String() {
		t.Errorf("err. actual: %v, expected: %v", i.String(), m.String())
	}

	i.Add(0)
	m.Add(0)
	if !checkHas(i, m) {
		t.Errorf("err. mismatch i & m")
	}
	if i.String() != m.String() {
		t.Errorf("err. actual: %v, expected: %v", i.String(), m.String())
	}

	i.Add(3)
	m.Add(3)
	if !checkHas(i, m) {
		t.Errorf("err. mismatch i & m")
	}
	if i.String() != m.String() {
		t.Errorf("err. actual: %v, expected: %v", i.String(), m.String())
	}
}

func Test_Union(t *testing.T) {
	i1 := IntSet{}
	i2 := IntSet{}

	m1 := newMapIntSet()
	m2 := newMapIntSet()

	i1.Add(4)
	m1.Add(4)

	i2.Add(2)
	m2.Add(2)

	i1.Add(1)
	m1.Add(1)

	i2.Add(1)
	m2.Add(1)

	i1.UnionWith(&i2)
	m1.UnionWith(m2)

	if i1.String() != m1.String() {
		t.Errorf("err. actual: %v, expected: %v", i1.String(), m1.String())
	}
}

func checkHas(i IntSet, m *mapIntset) bool {
	num := 0
	if i.Has(num) != m.Has(num) {
		return false
	}

	num = 1
	if i.Has(num) != m.Has(num) {
		return false
	}

	num = 3
	if i.Has(num) != m.Has(num) {
		return false
	}

	num = 100
	if i.Has(num) != m.Has(num) {
		return false
	}

	num = -1
	if i.Has(num) != m.Has(num) {
		return false
	}

	return true
}

type mapIntset struct {
	mp map[int]struct{}
}

func newMapIntSet() *mapIntset {
	m := &mapIntset{
		mp: make(map[int]struct{}),
	}
	return m
}

func (m *mapIntset) Add(i int) {
	if i >= 0 {
		m.mp[i] = struct{}{}
	}
}

func (m *mapIntset) Has(i int) bool {
	if i < 0 {
		return false
	}
	_, ok := m.mp[i]
	return ok
}

func (m *mapIntset) UnionWith(i *mapIntset) {
	for k := range i.mp {
		m.mp[k] = struct{}{}
	}
}

func (m *mapIntset) String() string {
	s := ""

	arr := []int{}
	for k := range m.mp {
		arr = append(arr, k)
	}
	sort.Ints(arr)

	for _, k := range arr {
		s += " " + strconv.Itoa(k)
	}
	s = "{" + strings.Trim(s, " ") + "}"
	return s
}
