package varcaser

// This file defines the CaseConvention type.

import (
	"strings"
	"unicode"
)

type WordCase func(string) string

// A CaseConvention is a way of writing variable names using separators and
// casing style.
type CaseConvention struct {
	JoinStyle
	SubsequentCase WordCase
	InitialCase    WordCase
	Example        string // Render the name of this case convention in itself
}

// A JoinStyle is a way of representing how individual components of a variable
// name are put together, and how to pull them apart.
type JoinStyle struct {
	Join  func([]string) string
	Split func(string) []string
}

// SimpleJoinStyle creates a JoinStyle that just splits and joins by a
// separator.
func SimpleJoinStyle(sep string) JoinStyle {
	return JoinStyle{
		Join: func(components []string) string {
			return strings.Join(components, sep)
		},
		Split: func(s string) []string {
			return strings.Split(s, sep)
		},
	}
}

// JoinStyle used in CamelCase. Special casing the Split function to keep
// acronynms together.
var camelJoinStyle = JoinStyle{
	Join: func(components []string) string {
		return strings.Join(components, "")

	},
	Split: func(s string) (components []string) {
		// NOTE(danver): While I keep finding new edge cases, I'll want
		// this to be easy-to-modify code rather than a regex.

		wasPreviousUpper := true
		current := []rune{}
		for _, c := range s {
			if wasPreviousUpper && unicode.IsUpper(c) {
				// If previous was uppercase, and this is
				// uppercase, continue the word.

				current = append(current, c)
			} else if wasPreviousUpper && !unicode.IsUpper(c) {

				// If the previous run was uppercase, but this
				// is not, set previous, but add it.

				// Edge case: the previous word was all uppercase.
				if len(current) > 1 {
					components = append(components, string(current[:len(current)-1]))
					current = current[len(current)-1:]
				}

				current = append(current, c)
				wasPreviousUpper = false
			} else if !wasPreviousUpper && unicode.IsUpper(c) {

				// If the previous rune was not uppercase, and
				// this character is, put current into
				// components first, then set wasPreviousUpper

				components = append(components, string(current))
				current = []rune{c}
				wasPreviousUpper = true
			} else if !wasPreviousUpper && !unicode.IsUpper(c) {

				// If the previous rune was not uppercase, and
				// this one is not, just add to this component.

				current = append(current, c)
			}
		}
		if len(current) != 0 {
			components = append(components, string(current))
		}
		return
	},
}

// SplitWords allows CaseConvention to implement Splitter.
func (c CaseConvention) SplitWords(s string) []string {
	return c.Split(s)
}

// ToStrictTitle returns the strict titling of a string without preserving
// existing caps in acronyms.
func ToStrictTitle(s string) string {
	return strings.Title(strings.ToLower(s))
}

// HttpAcronyms is effectively a set of acronyms that are conventionally
// uppercased in the HTTP Casing Convention.
var HttpAcronyms = map[string]bool{
	"XSS":  true,
	"SSL":  true,
	"HTTP": true,
	"MD5":  true,
	"TE":   true,
	"DNT":  true,
	"UIDH": true,
	"P3P":  true,
	"WWW":  true,
	"CSP":  true,
	"UA":   true,
}

// ToHttpTitle returns a string titled the way HTTP Headers title it.
func ToHttpTitle(s string) string {
	upper := strings.ToUpper(s)
	if _, ok := HttpAcronyms[upper]; ok {
		return upper
	}
	return ToStrictTitle(s)
}
