// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexElement
type AccountTypeBankCode struct {
	Value         *BankCodeType       `xml:",chardata"`
	International *w3c.Boolean        `xml:"International,attr,omitempty"`
	Prefix        *BankCodePrefixType `xml:"Prefix,attr,omitempty"`
}
