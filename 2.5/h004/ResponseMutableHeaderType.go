// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type ResponseMutableHeaderType struct {
	TransactionPhase *TransactionPhaseType                   `xml:"TransactionPhase"`
	SegmentNumber    *ResponseMutableHeaderTypeSegmentNumber `xml:"SegmentNumber,omitempty"`
	OrderID          *OrderIDType                            `xml:"OrderID,omitempty"`
	ReturnCode       *ReturnCodeType                         `xml:"ReturnCode"`
	ReportText       *ReportTextType                         `xml:"ReportText"`

	Any []*w3c.Any
}
