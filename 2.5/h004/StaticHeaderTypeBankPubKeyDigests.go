// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

// complex type
type StaticHeaderTypeBankPubKeyDigests struct {
	Authentication StaticHeaderTypeBankPubKeyDigestsAuthentication `xml:"Authentication"`
	Encryption     StaticHeaderTypeBankPubKeyDigestsEncryption     `xml:"Encryption"`
	Signature      StaticHeaderTypeBankPubKeyDigestsSignature      `xml:"Signature,omitempty"`
}
