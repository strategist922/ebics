// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import ds "github.com/openyard/ebics/3.0/xmldsig"
import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type PubKeyInfoType struct {
	X509Data *ds.X509Data `xml:"http://www.w3.org/2000/09/xmldsig# X509Data"`

	Any []*w3c.Any
}

func NewPubKeyInfoType() *PubKeyInfoType {
	return new(PubKeyInfoType)
}

func (me *PubKeyInfoType) SetX509Data(value *ds.X509Data) {
	me.X509Data = value
}
