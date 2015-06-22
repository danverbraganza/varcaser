// Package varcaser provides a way to change the case of variable names.
//
// TODO(danver): Flesh out these comments.
package varcaser

import "strings"

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
	JoinStyle:      camelJoinStyle,
	InitialCase:    strings.Title,
	SubsequentCase: strings.Title,
}

var LowerCamelCase = CaseConvention{
	JoinStyle:      camelJoinStyle,
	InitialCase:    strings.ToLower,
	SubsequentCase: strings.Title,
}
