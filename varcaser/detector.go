package varcaser

import (
	"fmt"
	"regexp"
	"unicode"
)

// ErrNoData is returned when an empty or nil slice is passed in.
// In this case, no useful information is returned from Detect().
var ErrNoData = fmt.Errorf("No data provided.")

// ErrNotEnoughData is returned when an accurate determination could not be
// made. However, the best guess will be provided in the returned CaseConvention.
var ErrNotEnoughData = fmt.Errorf("Not enough data provided.")

// ErrInconsistentData is returned when the input data contradicts itself. In
// this case, no useful information is returned from Detect().
var ErrInconsistentData = fmt.Errorf("Inconsistent data provided.")

type Detected struct {
	Split func(string) []string
}

func (d Detected) SplitWords(s string) []string {
	return d.Split(s)
}

// Detect returns a Splitter suitable for taking a variable name and turning it
// into a sequence of words.
func Detect(data []string) (sp Splitter, err error) {
	if len(data) == 0 {
		err = ErrNoData
		return
	}

	// joinSeparator can take the following values:
	// 0 -> assuming camelCase because not enough data
	// 1 -> camelCase, found at least one camelCase transition
	// any other punctuation rune -> that's the separator
	var joinSeparator rune = 0

	for _, s := range data {
		var innerErr error
		joinSeparator, innerErr = UpdateJoinStylePrediction(s, joinSeparator)
		if innerErr == ErrInconsistentData {
			err = ErrInconsistentData
			return
		}
	}

	d := Detected{}

	// We have now enough information to make a determination about the
	// separator.
	if joinSeparator <= 1 {
		d.Split = camelJoinStyle.Split
		if joinSeparator == 0 {
			err = ErrNotEnoughData
		}
	} else {
		d.Split = SimpleJoinStyle(string(joinSeparator)).Split
	}
	return d, err
}

// UpdateJoinStylePrediction works by comparing the data string to the
// rune that represents an expected JoinStyle, updating that rune as necessary.
func UpdateJoinStylePrediction(data string, sep rune) (rune, error) {
	punctsFound := map[rune]bool{}

	for _, r := range data {
		if unicode.IsPunct(r) {
			punctsFound[r] = true
		}
	}

	caseChangesFound, err := regexp.MatchString(`\p{Ll}\p{Lu}|\p{Lu}\p{Ll}`, data)
	if err != nil {
		print(err.Error())
		return -1, err
	}

	if len(punctsFound) == 0 && !caseChangesFound {
		// Could not find a symbol or a caseChange. Just give up and
		// continue with what we had before.
		return sep, nil
	}

	if len(punctsFound) > 1 {
		// Too many symbols found.
		return 0, ErrInconsistentData
	}

	if sep > 1 && !punctsFound[sep] {
		// sep was non-empty, which means we expected it to be found in this string.
		// Since we found at least one thing, we have an inconsistency.
		return 0, ErrInconsistentData
	}

	for k := range punctsFound {
		// Return the first and only symbol we found.
		return k, nil
	}

	if caseChangesFound {
		return 1, nil
	} else {
		return 0, nil
	}
}
