// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// Attribute
type Format struct {
	Value *w3c.Token `xml:"format,attr"`
}

func NewFormat() *Format {
	return new(Format)
}

func (me *Format) SetValue(value *w3c.Token) {
	me.Value = value
}
