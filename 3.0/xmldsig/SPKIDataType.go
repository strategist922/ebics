// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type SPKIDataType struct {
	SPKISexp *w3c.Base64Binary `xml:"SPKISexp"`

	Any []*w3c.Any
}
