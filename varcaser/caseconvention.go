package varcaser

// This file defines the CaseConvention type.

import (
	"strings"
	"unicode"
)

// A CaseConvention is a way of writing variable names using separators and
// casing style.
type CaseConvention struct {
	JoinStyle
	SubsequentCase func(string) string
	InitialCase    func(string) string
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

// ToStrictTitle returns the strict titling of a string without preserving
// existing caps in acronyms.
func ToStrictTitle(s string) string {
	return strings.Title(strings.ToLower(s))
}
