package varcaser

import (
	"fmt"
	"strings"
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

type WordCaseStatus int

type WordCaseStatusMap map[int]bool

// Detect tries to guess the casing convention by analysing a slice of strings.
// If the casing convention can be determined successfully, the resultant case
// convention is returned. Otherwise, an error explaining why the case detection
// could not be completed is returned.
func Detect(data []string) (c CaseConvention, err error) {
	if len(data) == 0 {
		err = ErrNoData
		return
	}

	initialCases := WordCaseStatusMap{}
	subsequentCases := WordCaseStatusMap{}
	joinSeparator := '\x00'

	for _, s := range data {
		var innerErr error
		joinSeparator, innerErr = UpdateJoinStylePrediction(s, joinSeparator)
		if innerErr == ErrInconsistentData {
			err = ErrInconsistentData
			return
		}
	}

	// We have now enough information to make a determination about the
	// separator.
	if joinSeparator == '\x00' {
		c.JoinStyle = camelJoinStyle
	} else {
		c.JoinStyle = SimpleJoinStyle(string(joinSeparator))
	}

	for _, s := range data {
		for i, word := range c.JoinStyle.Split(s) {
			if i == 0 {
				initialCases.UpdateCasePrediction(word)
			} else {
				subsequentCases.UpdateCasePrediction(word)
			}
		}
	}

	c.InitialCase, err = initialCases.GetPrediction(joinSeparator == '\x00')

	if err != nil {
		return
	}
	c.SubsequentCase, err = subsequentCases.GetPrediction(joinSeparator == '\x00')

	return
}

// allPossibleCases is an ordered, intended-to-be-immutable, list of the cases
// we expect our data to come from. The order here influences precedenc, with
// earlier entries being preferred. In case we have insufficient data to detect
// between StrictTitle and HttpTitle, for example, we'll prefer to assume
// StrictTitle
var allPossibleCases = []WordCase{
	strings.ToLower, // 0
	strings.ToUpper, // 1
	ToStrictTitle,   // 2
	strings.Title,   // 3
	ToHttpTitle,     // 4

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

	if len(punctsFound) == 0 {
		// Could not find a symbol. Just give up and continue with what
		// we had before.
		return sep, nil
	}

	if len(punctsFound) > 1 {
		// Too many symbols found.
		return '\x00', ErrInconsistentData
	}

	if sep != '\x00' && !punctsFound[sep] {
		// sep was non-empty, which means we expected it to be found in this string.
		// Since we found at least one thing, we have an inconsistency.
		return '\x00', ErrInconsistentData
	}

	for k := range punctsFound {
		// Return the first and only symbol we found.
		return k, nil
	}

	// We should never get here.
	return '\x00', ErrNotEnoughData
}

// UpdateCasePrediction takes a list of valid case guesses, and a word string,
// and notes all the guesses that matched that string.
func (wcs WordCaseStatusMap) UpdateCasePrediction(s string) {
	if len(s) == 0 {
		return
	}

	if len(s) == 1 {
		// We can't tell if it's a Title, Upper, or any other case. Only
		// if it's lower.
		if allPossibleCases[0](s) == s {
			wcs[0] = true
		}
		return
	}

	for i, c := range allPossibleCases {
		if c(s) == s {
			wcs[i] = true
		}
	}
}

// GetPrediction returns the best prediction, if possible, for a given sequence
// of seen word cases.
func (wcs WordCaseStatusMap) GetPrediction(isCamelCase bool) (WordCase, error) {
	if len(wcs) == 0 {
		// We didn't see any cases.
		return nil, ErrNotEnoughData
	}

	if len(wcs) == 1 {
		// We see exactly one case, so return it.
		for i, _ := range wcs {
			return allPossibleCases[i], nil
		}
	}

	delete(wcs, 4) // HttpTitle could only win if it's the only case left
	// standing. It didn't win, so get rid of it.

	// Special case: camelcase can cause confusion between TitleCase,
	// StrictTitleCase, HttpTitleCase and UpperCase
	if isCamelCase && wcs[3] { // TitleCase is in play.
		if wcs[1] { // There was an uppercase string.
			delete(wcs, 1) // Remove uppercase
			delete(wcs, 2) // Remove strict title
		} else {
			delete(wcs, 3) // Remove lax title.
		}
	} else if wcs[3] { // No CamelCase, so just get rid of duped
			   // StrictTitle/TitleCase.
		delete(wcs, 3)
	}


	if len(wcs) == 1 {
		// We saw exactly one case, so return it.
		for i, _ := range wcs {
			return allPossibleCases[i], nil
		}
	}

	return nil, ErrInconsistentData
}
