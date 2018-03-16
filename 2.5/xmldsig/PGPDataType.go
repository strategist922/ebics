// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type PGPDataType struct {
	PGPKeyID     *w3c.Base64Binary `xml:"PGPKeyID"`
	PGPKeyPacket *w3c.Base64Binary `xml:"PGPKeyPacket,omitempty"`

	Any []*w3c.Any
}

func NewPGPDataType() *PGPDataType {
	return new(PGPDataType)
}

func (me *PGPDataType) SetPGPKeyID(value *w3c.Base64Binary) {
	me.PGPKeyID = value
}

func (me *PGPDataType) AddPGPKeyID() *w3c.Base64Binary {
	me.PGPKeyID = new(w3c.Base64Binary)
	return me.PGPKeyID
}

func (me *PGPDataType) SetPGPKeyPacket(value *w3c.Base64Binary) {
	me.PGPKeyPacket = value
}

func (me *PGPDataType) AddPGPKeyPacket() *w3c.Base64Binary {
	me.PGPKeyPacket = new(w3c.Base64Binary)
	return me.PGPKeyPacket
}
