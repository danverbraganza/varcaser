package varcaser

import (
	"reflect"
	"testing"
)

func AssertEqual(specimen, expected interface{}, t *testing.T) {
	if !reflect.DeepEqual(specimen, expected) {
		t.Errorf("Wanted %v, got %v", expected, specimen)
	}

}

func TestCamelSplit(t *testing.T) {
	expected := []string{"Test", "Camel", "Split"}
	specimen := camelJoinStyle.Split("TestCamelSplit")
	AssertEqual(specimen, expected, t)
}

func TestCamelSplitComplex(t *testing.T) {
	expected := []string{"async", "HTTP", "Router4"}
	specimen := camelJoinStyle.Split("asyncHTTPRouter4")
	AssertEqual(specimen, expected, t)

}

func TestCamelSplitMixedUp(t *testing.T) {
	specimen := camelJoinStyle.Split("FooBarMVCFrameworkHttp")
	expected := []string{"Foo", "Bar", "MVC", "Framework", "Http"}
	AssertEqual(specimen, expected, t)
}

func TestCaserSimple(t *testing.T) {
	c := Caser{From: LowerSnakeCase, To: UpperCamelCase}

	specimen := c.String("my_int_var_20")
	expected := "MyIntVar20"
	AssertEqual(specimen, expected, t)
}

func TestCaserLeadingUnderscoreToCamel(t *testing.T) {
	// This was a tricky case. I expected that _private_method would convert
	// to _PrivateMethod, but to get that behaviour complicates the
	// conversion functions (or priviliges _ as a conversion), and it seems
	// that
	// http://www.oracle.com/technetwork/java/javase/documentation/codeconventions-135099.html#367
	// doesn't approve of leading underscores in that case anyway.

	// This could be supported, but for the nonce I'm going to assume YAGNI.

	c := Caser{From: LowerSnakeCase, To: UpperCamelCase}

	specimen := c.String("_private_method")
	expected := "PrivateMethod" // NOT "_PrivateMethod"
	AssertEqual(specimen, expected, t)
}

func TestCaserLeadingUnderscoreToSnake(t *testing.T) {
	c := Caser{From: LowerSnakeCase, To: ScreamingSnakeCase}

	specimen := c.String("_private_method")
	expected := "_PRIVATE_METHOD"
	AssertEqual(specimen, expected, t)
}

func TestCaserCamelToKebab(t *testing.T) {
	c := Caser{From: UpperCamelCase, To: KebabCase}

	specimen := c.String("SomeInitMethod")
	expected := "some-init-method"
	AssertEqual(specimen, expected, t)
}

func TestCaserSnakeToUpperKebab(t *testing.T) {
	c := Caser{From: LowerSnakeCase, To: UpperKebabCase}

	specimen := c.String("some_init_method")
	expected := "Some-Init-Method"
	AssertEqual(specimen, expected, t)
}

func TestCaserSnakeToScreamingKebab(t *testing.T) {
	c := Caser{From: LowerSnakeCase, To: ScreamingKebabCase}

	specimen := c.String("some_init_method")
	expected := "SOME-INIT-METHOD"
	AssertEqual(specimen, expected, t)
}

// AsyncHTTPRequest -> AsyncHttpRequest
func TestCaserCamelToCamelLoseCapitals(t *testing.T) {
	c := Caser{From: UpperCamelCase, To: UpperCamelCase}

	specimen := c.String("AsyncHTTPRequest")
	expected := "AsyncHttpRequest"
	AssertEqual(specimen, expected, t)
}

// AsyncHTTPRequest -> AsyncHttpRequest
func TestCaserCamelToCamelKeepCapitals(t *testing.T) {
	c := Caser{From: UpperCamelCase, To: UpperCamelCaseKeepCaps}

	specimen := c.String("AsyncHTTPRequest")
	expected := "AsyncHTTPRequest"
	AssertEqual(specimen, expected, t)
}

func TestCaserLowerCamelInitialCapital(t *testing.T) {
	// This is another tricky case. I decided that the initial Capital does
	// NOT indicate a hidden initial separator, but that might change.

	c := Caser{From: LowerCamelCase, To: KebabCase}
	specimen := c.String("SomeInitMethod")
	expected := "some-init-method" // NOT "-some-init-method"
	AssertEqual(specimen, expected, t)
}
