// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type MutableHeaderType struct {
	TransactionPhase *TransactionPhaseType           `xml:"TransactionPhase"`
	SegmentNumber    *MutableHeaderTypeSegmentNumber `xml:"SegmentNumber,omitempty"`

	Any []*w3c.Any
}

func NewMutableHeaderType() *MutableHeaderType {
	return new(MutableHeaderType)
}

func (me *MutableHeaderType) SetTransactionPhase(value *TransactionPhaseType) {
	me.TransactionPhase = value
}

func (me *MutableHeaderType) AddTransactionPhase() *TransactionPhaseType {
	me.TransactionPhase = new(TransactionPhaseType)
	return me.TransactionPhase
}

func (me *MutableHeaderType) SetSegmentNumber(value *MutableHeaderTypeSegmentNumber) {
	me.SegmentNumber = value
}

func (me *MutableHeaderType) AddSegmentNumber() *MutableHeaderTypeSegmentNumber {
	me.SegmentNumber = new(MutableHeaderTypeSegmentNumber)
	return me.SegmentNumber
}
