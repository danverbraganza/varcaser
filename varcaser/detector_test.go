package varcaser

import (
	"reflect"
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
	AssertEqual(c.SplitWords("a_b"), []string{"a_b"}, t)
}

func TestDetectorUnderscoreLowercase(t *testing.T) {
	c, err := Detect([]string{"abcd", "_myvar", "this_is_my_var"})
	AssertEqual(err, nil, t)
	AssertEqual(c.SplitWords("a_b"), []string{"a", "b"}, t)
}

func TestDetectorCamelTitleCase(t *testing.T) {
	c, err := Detect([]string{"Abcd", "MyVar", "ThisIsMyVar"})
	AssertEqual(err, nil, t)
	AssertEqual(c.SplitWords("a_b"), []string{"a_b"}, t)
	AssertEqual(c.SplitWords("AbcdDef"), []string{"Abcd", "Def"}, t)
}

func TestDetectorHyphenLowerTitleCase(t *testing.T) {
	c, err := Detect([]string{"a-B", "my", "my-Big-Variable"})
	AssertEqual(err, nil, t)
	AssertEqual(c.SplitWords("a_b"), []string{"a_b"}, t)
	AssertEqual(c.SplitWords("a_B"), []string{"a_B"}, t)
	AssertEqual(c.SplitWords("a-B-c"), []string{"a", "B", "c"}, t)
}
