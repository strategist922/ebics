// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package s002

import w3c "github.com/openyard/ebics/3.0/w3c"

// SimpleType
type UserIDType w3c.Token

func NewUserIDType(value *w3c.Token) *UserIDType {
	me := (*UserIDType)(value)
	return me
}
