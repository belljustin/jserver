package hashset

import (
	"regexp"
)

// OrCaseInsensitiveRegexp returns a Regexp object that matches any of the words
// in the case-insensitive set.
// e.g. (?i)(GET|POST) would match Get or pOsT, etc.
func (s *StringSet) OrCaseInsensitiveRegexp() *regexp.Regexp {
	re := regexp.MustCompile(",")
	return regexp.MustCompile("(?i)(" +
		re.ReplaceAllString(s.String(), "|") +
		")")
}
