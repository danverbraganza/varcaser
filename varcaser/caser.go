package varcaser

import (
	"golang.org/x/text/transform"
)

// type Caser is a text transformer that takes converts a variable from one
// casing convention to another.
type Caser struct {
	From CaseConvention
	To   CaseConvention
	transform.NopResetter
}

// String returns the representation of a variable name in this Caser's To
// CaseConvention given a variable name in this Caser's From CaseConvention.
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

// Bytes is provided for compatibility with the Transformer interface. Since
// Caser has no special treatement of bytes, the bytes are converted to and from
// strings. Warning: Untested.
func (c Caser) Bytes(b []byte) (result []byte) {
	copy(result, c.String(string(b)))
	return
}

// Provided for compatibility with the Transformer interface. Since Caser has no
// special treatement of bytes, the bytes are converted to and from strings.
// Warning: Untested.
func (c Caser) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	nSrc = len(src) // Always read all the bytes of src
	result := c.Bytes(src)
	if len(result) > cap(dst) {
		err = transform.ErrShortDst
	}
	nDst = copy(dst, src)
	return
}
