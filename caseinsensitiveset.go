package caseinsensitiveset

import (
	mapset "github.com/deckarep/golang-set/v2"
	"strings"
)

type CaseInsensitiveSet struct {
	s          mapset.Set[string]
	elementMap map[string]string
}

func NewCaseInsensitiveSet(values ...string) *CaseInsensitiveSet {
	set := CaseInsensitiveSet{
		s:          mapset.NewSet[string](),
		elementMap: make(map[string]string),
	}

	for _, v := range values {
		set.elementMap[strings.ToLower(v)] = v
		set.s.Add(strings.ToLower(v))
	}

	return &set
}

func (s *CaseInsensitiveSet) Add(value string) bool {
	s.elementMap[strings.ToLower(value)] = value
	return s.s.Add(strings.ToLower(value))
}

func (s *CaseInsensitiveSet) Remove(value string) {
	delete(s.elementMap, strings.ToLower(value))
	s.s.Remove(strings.ToLower(value))
}

func (s *CaseInsensitiveSet) Difference(s1 *CaseInsensitiveSet) *CaseInsensitiveSet {
	diffSet := NewCaseInsensitiveSet()
	for _, v := range s.s.Difference(s1.s).ToSlice() {
		diffSet.Add(s.elementMap[v])
	}
	return diffSet
}

func (s *CaseInsensitiveSet) ToSlice() []string {
	vals := make([]string, 0, len(s.elementMap))

	for _, v := range s.elementMap {
		vals = append(vals, v)
	}
	return vals
}

func (s *CaseInsensitiveSet) Contains(value ...string) bool {
	return s.s.Contains(mapStringSlice(value, strings.ToLower)...)
}

func mapStringSlice(t []string, f func(string) string) []string {
	s := make([]string, len(t))
	for i, v := range t {
		s[i] = f(v)
	}
	return s
}
