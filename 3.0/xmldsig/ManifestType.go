// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type ManifestType struct {
	Id        w3c.ID    `xml:"Id,attr,omitempty"`
	Reference Reference `xml:"Reference"`
}