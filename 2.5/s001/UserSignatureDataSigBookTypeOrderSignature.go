// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package s001

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex element
type UserSignatureDataSigBookTypeOrderSignature struct {
	Value     *OrderSignatureType `xml:",chardata"`
	PartnerID *w3c.String         `xml:"PartnerID,attr"`
}
