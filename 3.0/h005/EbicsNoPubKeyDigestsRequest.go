// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

// complex type
type EbicsNoPubKeyDigestsRequest struct {
	Header        EbicsNoPubKeyDigestsRequestHeader `xml:"header"`
	Body          EbicsNoPubKeyDigestsRequestBody   `xml:"body"`
	AuthSignature AuthSignature                     `xml:"AuthSignature"`
	VersionAttrGroup
}