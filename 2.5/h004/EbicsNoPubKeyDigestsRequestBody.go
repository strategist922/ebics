// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h004

import ds "github.com/openyard/ebics/2.5/xmldsig"

// complex type
type EbicsNoPubKeyDigestsRequestBody struct {
	X509Data *ds.X509Data `xml:"http://www.w3.org/2000/09/xmldsig# X509Data,omitempty"`
}
