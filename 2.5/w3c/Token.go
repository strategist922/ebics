// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package w3c

// BaseType
type Token string

func NewToken(value string) *Token {
	me := (*Token)(&value)
	return me
}
