// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// Attribute
type Id struct {
	Value *w3c.ID `xml:"Id,attr"`
}

func NewId() *Id {
	return new(Id)
}

func (me *Id) SetValue(value *w3c.ID) {
	me.Value = value
}
