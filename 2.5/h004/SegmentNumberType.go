// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// SimpleType
type SegmentNumberType w3c.PositiveInteger

func NewSegmentNumberType(value *w3c.PositiveInteger) *SegmentNumberType {
	me := (*SegmentNumberType)(value)
	return me
}
