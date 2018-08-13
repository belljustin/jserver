package hashset

import (
	"strings"
)

type StringSet struct {
	set map[string]struct{}
}

func (s *StringSet) Add(value string) bool {
	_, found := s.set[value]
	s.set[value] = *new(struct{})
	return !found
}

func NewStringSet(values ...string) *StringSet {
	s := new(StringSet)
	s.set = make(map[string]struct{})

	for _, v := range values {
		s.Add(v)
	}
	return s
}

func (s *StringSet) Contains(value string) bool {
	_, found := s.set[value]
	return found
}

func (s *StringSet) String() string {
	var b strings.Builder

	first := true
	for k, _ := range s.set {
		if !first { // add commas between strings
			b.WriteByte(',')
		} else {
			first = false
		}

		b.WriteString(k)
	}

	return b.String()
}
