package varcaser

import (
	"golang.org/x/text/transform"
)

// This file defines the Caser object, which perfoms most of the case
// conversions.

type Caser struct {
	From CaseConvention
	To CaseConvention
	transform.NopResetter
}

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

// WARNING: The following methods have not been tested yet.

// Provided for compatibility with Caser interface. No special treatement of
// bytes.
func (c Caser) Bytes(b []byte) (result []byte) {
	copy(result, c.String(string(b)))
	return
}

func (c Caser) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	nSrc = len(src) // Always read all the bytes of src
	result := c.Bytes(src)
	if len(result) > cap(dst) {
		err = transform.ErrShortDst
	}
	nDst = copy(dst, src)
	return
}
