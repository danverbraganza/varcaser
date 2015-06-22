package varcaser

import "errors"

// This file defines the Caser object, which perfoms most of the case
// conversions.

type Caser struct {
	From CaseConvention
	To CaseConvention
}

func (c Caser) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	err = errors.New("Not implemented")
	return
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
