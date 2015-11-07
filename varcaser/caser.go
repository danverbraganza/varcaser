package varcaser

import (
	"golang.org/x/text/transform"
)

// type Caser is a text transformer that takes converts a variable from one
// casing convention to another.
type Caser struct {
	From Splitter
	To   CaseConvention
	transform.NopResetter
}

// Splitter is an interface for a type that can decompose a variable name into
// its component words.
type Splitter interface {
	SplitWords(string) []string
}

// String returns the representation of a variable name in this Caser's To
// CaseConvention given a variable name in this Caser's From CaseConvention.
func (c Caser) String(s string) string {
	components := []string{}
	for i, s := range c.From.SplitWords(s) {
		if i == 0 {
			components = append(components, c.To.InitialCase(s))
		} else {
			components = append(components, c.To.SubsequentCase(s))
		}
	}
	return c.To.Join(components)
}

// Bytes is provided for compatibility with the Transformer interface. Since
// Caser has no special treatement of bytes, the bytes are converted to and from
// strings.
func (c Caser) Bytes(b []byte) []byte {
	return []byte(c.String(string(b)))
}

// Provided for compatibility with the Transformer interface. Since Caser has no
// special treatement of bytes, the bytes are converted to and from strings.
// Will treat the entirity of src as ONE variable name.
func (c Caser) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	nSrc = len(src) // Always read all the bytes of src
	result := c.Bytes(src)
	if len(result) > len(dst) {
		err = transform.ErrShortDst
	}
	nDst = copy(dst, result)
	return
}
