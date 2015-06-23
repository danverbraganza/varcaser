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
	c := Caser{From: lower_snake_case, To: UpperCamelCase}

	specimen := c.String("my_int_var_20")
	expected := "MyIntVar20"
	AssertEqual(specimen, expected, t)
}
