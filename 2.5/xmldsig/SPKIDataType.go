// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type SPKIDataType struct {
	SPKISexp w3c.Base64Binary `xml:"SPKISexp"`

	Any []*w3c.Any
}
