// Package varcaser provides a way to change the case of variable names.

package varcaser

import (
	"errors"
	"strings"
	"unicode"
)

// A CaseConvention is a way of writing variable names using titling.
type CaseConvention struct {
	JoinStyle
	SubsequentCase func(string) string
	InitialCase    func(string) string
}

type JoinStyle struct {
	Join  func([]string) string
	Split func(string) []string
}

// A simple join style can be specified by a string separator.
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

var camelJoinStyle = JoinStyle{
	Join: func(components []string) string {
		return strings.Join(components, "")

	},
	Split: func(s string) (components []string) {
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
					components = append(components, string(current[:len(current) - 1]))
					current = current[len(current) - 1:]
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

var lower_snake_case = CaseConvention{
	JoinStyle:      SimpleJoinStyle("_"),
	InitialCase:    strings.ToLower,
	SubsequentCase: strings.ToLower,
}

var UPPER_SNAKE_CASE = CaseConvention{
	JoinStyle:      SimpleJoinStyle("_"),
	InitialCase:    strings.ToUpper,
	SubsequentCase: strings.ToUpper,
}

var UpperCamelCase = CaseConvention{
	JoinStyle: camelJoinStyle,
	InitialCase: strings.Title,
	SubsequentCase: strings.Title,
}

var LowerCamelCase = CaseConvention{
	JoinStyle: camelJoinStyle,
	InitialCase: strings.ToLower,
	SubsequentCase: strings.Title,
}


type Caser struct {
	From CaseConvention
	To CaseConvention
}

func (c Caser) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	err = errors.New("Not implemented")
	return
}

func (c Caser) String(s string) string {
	components := []string{}
	for i, s := range c.From.Split(s) {
		if i == 0 {
			components = append(components, c.To.InitialCase(s))
		} else {
			components = append(components, c.To.SubsequentCase(s))
		}
	}
	return c.To.Join(components)
}
