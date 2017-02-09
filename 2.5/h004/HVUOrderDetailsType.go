// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HVUOrderDetailsType struct {
	OrderType      OrderTBaseType        `xml:"OrderType"`
	FileFormat     FileFormatType        `xml:"FileFormat,omitempty"`
	OrderID        OrderIDType           `xml:"OrderID"`
	OrderDataSize  w3c.PositiveInteger   `xml:"OrderDataSize"`
	SigningInfo    HVUSigningInfoType    `xml:"SigningInfo"`
	SignerInfo     SignerInfoType        `xml:"SignerInfo,omitempty"`
	OriginatorInfo HVUOriginatorInfoType `xml:"OriginatorInfo"`

	Any []*w3c.Any
}
