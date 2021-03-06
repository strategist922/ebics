// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type X509DataType struct {
	X509IssuerSerial *X509IssuerSerialType `xml:"X509IssuerSerial"`
	X509SKI          *w3c.Base64Binary     `xml:"X509SKI"`
	X509SubjectName  *w3c.String           `xml:"X509SubjectName"`
	X509Certificate  *w3c.Base64Binary     `xml:"X509Certificate"`
	X509CRL          *w3c.Base64Binary     `xml:"X509CRL"`
}

func NewX509DataType() *X509DataType {
	return new(X509DataType)
}

func (me *X509DataType) SetX509IssuerSerial(value *X509IssuerSerialType) {
	me.X509IssuerSerial = value
}

func (me *X509DataType) AddX509IssuerSerial() *X509IssuerSerialType {
	me.X509IssuerSerial = new(X509IssuerSerialType)
	return me.X509IssuerSerial
}

func (me *X509DataType) SetX509SKI(value *w3c.Base64Binary) {
	me.X509SKI = value
}

func (me *X509DataType) AddX509SKI() *w3c.Base64Binary {
	me.X509SKI = new(w3c.Base64Binary)
	return me.X509SKI
}

func (me *X509DataType) SetX509SubjectName(value *w3c.String) {
	me.X509SubjectName = value
}

func (me *X509DataType) AddX509SubjectName() *w3c.String {
	me.X509SubjectName = new(w3c.String)
	return me.X509SubjectName
}

func (me *X509DataType) SetX509Certificate(value *w3c.Base64Binary) {
	me.X509Certificate = value
}

func (me *X509DataType) AddX509Certificate() *w3c.Base64Binary {
	me.X509Certificate = new(w3c.Base64Binary)
	return me.X509Certificate
}

func (me *X509DataType) SetX509CRL(value *w3c.Base64Binary) {
	me.X509CRL = value
}

func (me *X509DataType) AddX509CRL() *w3c.Base64Binary {
	me.X509CRL = new(w3c.Base64Binary)
	return me.X509CRL
}
