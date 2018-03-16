// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// AttributeGroup
type AuthenticationMarker struct {
	Authenticate *w3c.Boolean `xml:"authenticate,attr"`
}

func NewAuthenticationMarker() *AuthenticationMarker {
	return new(AuthenticationMarker)
}

func (me *AuthenticationMarker) SetAuthenticate(value *w3c.Boolean) {
	me.Authenticate = value
}
