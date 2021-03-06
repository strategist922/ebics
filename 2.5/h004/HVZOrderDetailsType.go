// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type HVZOrderDetailsType struct {
	OrderType             *OrderTBaseType        `xml:"OrderType"`
	FileFormat            *FileFormatType        `xml:"FileFormat,omitempty"`
	OrderID               *OrderIDType           `xml:"OrderID"`
	DataDigest            *DataDigestType        `xml:"DataDigest"`
	OrderDataAvailable    *w3c.Boolean           `xml:"OrderDataAvailable"`
	OrderDataSize         *w3c.PositiveInteger   `xml:"OrderDataSize"`
	OrderDetailsAvailable *w3c.Boolean           `xml:"OrderDetailsAvailable"`
	SigningInfo           *HVUSigningInfoType    `xml:"SigningInfo"`
	SignerInfo            []*SignerInfoType      `xml:"SignerInfo,omitempty"`
	OriginatorInfo        *HVUOriginatorInfoType `xml:"OriginatorInfo"`

	Any []*w3c.Any
}

func NewHVZOrderDetailsType() *HVZOrderDetailsType {
	return new(HVZOrderDetailsType)
}

func (me *HVZOrderDetailsType) SetOrderType(value *OrderTBaseType) {
	me.OrderType = value
}

func (me *HVZOrderDetailsType) AddOrderType() *OrderTBaseType {
	me.OrderType = new(OrderTBaseType)
	return me.OrderType
}

func (me *HVZOrderDetailsType) SetFileFormat(value *FileFormatType) {
	me.FileFormat = value
}

func (me *HVZOrderDetailsType) AddFileFormat() *FileFormatType {
	me.FileFormat = new(FileFormatType)
	return me.FileFormat
}

func (me *HVZOrderDetailsType) SetOrderID(value *OrderIDType) {
	me.OrderID = value
}

func (me *HVZOrderDetailsType) AddOrderID() *OrderIDType {
	me.OrderID = new(OrderIDType)
	return me.OrderID
}

func (me *HVZOrderDetailsType) SetDataDigest(value *DataDigestType) {
	me.DataDigest = value
}

func (me *HVZOrderDetailsType) AddDataDigest() *DataDigestType {
	me.DataDigest = new(DataDigestType)
	return me.DataDigest
}

func (me *HVZOrderDetailsType) SetOrderDataAvailable(value *w3c.Boolean) {
	me.OrderDataAvailable = value
}

func (me *HVZOrderDetailsType) AddOrderDataAvailable() *w3c.Boolean {
	me.OrderDataAvailable = new(w3c.Boolean)
	return me.OrderDataAvailable
}

func (me *HVZOrderDetailsType) SetOrderDataSize(value *w3c.PositiveInteger) {
	me.OrderDataSize = value
}

func (me *HVZOrderDetailsType) AddOrderDataSize() *w3c.PositiveInteger {
	me.OrderDataSize = new(w3c.PositiveInteger)
	return me.OrderDataSize
}

func (me *HVZOrderDetailsType) SetOrderDetailsAvailable(value *w3c.Boolean) {
	me.OrderDetailsAvailable = value
}

func (me *HVZOrderDetailsType) AddOrderDetailsAvailable() *w3c.Boolean {
	me.OrderDetailsAvailable = new(w3c.Boolean)
	return me.OrderDetailsAvailable
}

func (me *HVZOrderDetailsType) SetSigningInfo(value *HVUSigningInfoType) {
	me.SigningInfo = value
}

func (me *HVZOrderDetailsType) AddSigningInfo() *HVUSigningInfoType {
	me.SigningInfo = new(HVUSigningInfoType)
	return me.SigningInfo
}

func (me *HVZOrderDetailsType) AddSignerInfo(value *SignerInfoType) {
	me.SignerInfo = append(me.SignerInfo, value)
}

func (me *HVZOrderDetailsType) SetOriginatorInfo(value *HVUOriginatorInfoType) {
	me.OriginatorInfo = value
}

func (me *HVZOrderDetailsType) AddOriginatorInfo() *HVUOriginatorInfoType {
	me.OriginatorInfo = new(HVUOriginatorInfoType)
	return me.OriginatorInfo
}
