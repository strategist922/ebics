// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

// complex type
type H3KRequestOrderDataType struct {
	SignatureCertificateInfo      SignatureCertificateInfoType      `xml:"SignatureCertificateInfo"`
	AuthenticationCertificateInfo AuthenticationCertificateInfoType `xml:"AuthenticationCertificateInfo"`
	EncryptionCertificateInfo     EncryptionCertificateInfoType     `xml:"EncryptionCertificateInfo"`
	PartnerID                     PartnerIDType                     `xml:"PartnerID"`
	UserID                        UserIDType                        `xml:"UserID"`
}
