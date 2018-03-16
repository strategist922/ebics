// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// Attribute
type URI struct {
	Value *w3c.AnyURI `xml:"URI,attr"`
}

func NewURI() *URI {
	return new(URI)
}

func (me *URI) SetValue(value *w3c.AnyURI) {
	me.Value = value
}
