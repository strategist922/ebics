// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexElement
type MutableHeaderTypeSegmentNumber struct {
	Value       *SegmentNumberType `xml:",chardata"`
	LastSegment *w3c.Boolean       `xml:"LastSegment,attr"`
}

func NewMutableHeaderTypeSegmentNumber() *MutableHeaderTypeSegmentNumber {
	return new(MutableHeaderTypeSegmentNumber)
}

func (me *MutableHeaderTypeSegmentNumber) SetLastSegment(value *w3c.Boolean) {
	me.LastSegment = value
}
