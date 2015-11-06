package varcaser

import (
	"reflect"
	"strings"
	"testing"
)

func AssertIdentical(actual interface{}, expected interface{}, t *testing.T) {
	AssertEqual(
		reflect.ValueOf(actual).Pointer(),
		reflect.ValueOf(expected).Pointer(), t)

}

func TestDetectorEmpty(t *testing.T) {
	_, err := Detect(nil)
	AssertEqual(err, ErrNoData, t)
	// The returned CaseConvention could be anything.
}

func TestDetectorVeryAmbiguous(t *testing.T) {
	c, err := Detect([]string{"abcd"})
	AssertEqual(err, ErrNotEnoughData, t)
	AssertIdentical(c.JoinStyle.Split, camelJoinStyle.Split, t)
	AssertIdentical(c.InitialCase, strings.ToLower, t)
}

func TestDetectorUnderscoreLowercase(t *testing.T) {
	c, err := Detect([]string{"abcd", "_myvar", "this_is_my_var"})
	AssertEqual(err, nil, t)
	AssertEqual(c.JoinStyle.Join([]string{"a", "b"}), "a_b", t)
	AssertIdentical(c.InitialCase, strings.ToLower, t)
	AssertIdentical(c.SubsequentCase, strings.ToLower, t)
}

func TestDetectorCamelTitleCase(t *testing.T) {
	c, err := Detect([]string{"Abcd", "MyVar", "ThisIsMyVar"})
	AssertEqual(err, nil, t)
	AssertEqual(c.JoinStyle.Join([]string{"a", "b"}), "ab", t)
	AssertIdentical(c.InitialCase, ToStrictTitle, t)
	AssertIdentical(c.SubsequentCase, ToStrictTitle, t)
}

func TestDetectorCamelLaxCase(t *testing.T) {
	c, err := Detect([]string{"MyHTTPString", "NoVariable", "T"})
	AssertEqual(err, nil, t)
	AssertEqual(c.JoinStyle.Join([]string{"a", "b"}), "ab", t)
	AssertEqual(c.InitialCase("abcd"), "Abcd", t)
	AssertEqual(c.InitialCase("ABCD"), "Abcd", t)

	AssertEqual(c.SubsequentCase("abcd"), "Abcd", t)
	AssertEqual(c.SubsequentCase("ABCD"), "ABCD", t)
}

func TestDetectorHyphenLowerTitleCase(t *testing.T) {
	c, err := Detect([]string{"a-B", "my", "my-Big-Variable"})
	AssertEqual(err, nil, t)
	AssertEqual(c.JoinStyle.Join([]string{"a", "b"}), "a-b", t)
	AssertEqual(c.InitialCase("abcd"), "abcd", t)
	AssertEqual(c.InitialCase("ABCD"), "abcd", t)

	AssertEqual(c.SubsequentCase("abcd"), "Abcd", t)
	AssertEqual(c.SubsequentCase("ABCD"), "Abcd", t)

}
