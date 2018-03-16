// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type HVSRequestOrderDataType struct {
	CancelledDataDigest *DataDigestType `xml:"CancelledDataDigest"`

	Any []*w3c.Any
}

func NewHVSRequestOrderDataType() *HVSRequestOrderDataType {
	return new(HVSRequestOrderDataType)
}

func (me *HVSRequestOrderDataType) SetCancelledDataDigest(value *DataDigestType) {
	me.CancelledDataDigest = value
}

func (me *HVSRequestOrderDataType) AddCancelledDataDigest() *DataDigestType {
	me.CancelledDataDigest = new(DataDigestType)
	return me.CancelledDataDigest
}
