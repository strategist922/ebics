// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type CanonicalizationMethodType struct {
	Algorithm w3c.AnyURI `xml:"Algorithm,attr"`

	Any []*w3c.Any
}