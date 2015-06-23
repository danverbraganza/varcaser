// Package varcaser provides a way to change the case of variable names.
//
// TODO(danver): Flesh out these comments.
package varcaser

import "strings"

var LowerSnakeCase = CaseConvention{
	JoinStyle:      SimpleJoinStyle("_"),
	InitialCase:    strings.ToLower,
	SubsequentCase: strings.ToLower,
	Example: "lower_snake_case",
}

var ScreamingSnakeCase = CaseConvention{
	JoinStyle:      SimpleJoinStyle("_"),
	InitialCase:    strings.ToUpper,
	SubsequentCase: strings.ToUpper,
	Example: "LOWER_SNAKE_CASE",
}

var KebabCase = CaseConvention{
	JoinStyle:      SimpleJoinStyle("-"),
	InitialCase:    strings.ToLower,
	SubsequentCase: strings.ToLower,
	Example: "kebab-case",
}

var UpperKebabCase = CaseConvention{
	JoinStyle:      SimpleJoinStyle("-"),
	InitialCase:    strings.ToUpper,
	SubsequentCase: strings.ToUpper,
	Example: "Upper-Kebab-Case",
}

var SpinalCase = KebabCase // Hate this name, but some people use it.
var TrainCase = UpperKebabCase // Ditto.

var UpperCamelCase = CaseConvention{
	JoinStyle:      camelJoinStyle,
	InitialCase:    strings.Title,
	SubsequentCase: strings.Title,
	Example: "UpperCamelCase",
}

var LowerCamelCase = CaseConvention{
	JoinStyle:      camelJoinStyle,
	InitialCase:    strings.ToLower,
	SubsequentCase: strings.Title,
	Example: "lowerCamelCase",
}
