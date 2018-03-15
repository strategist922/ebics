// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type KeyValueType struct {
	DSAKeyValue *DSAKeyValue `xml:"DSAKeyValue"`
	RSAKeyValue *RSAKeyValue `xml:"RSAKeyValue"`

	Any []*w3c.Any
}
