// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

// ComplexType
type StandardOrderParamsType struct {
	DateRange *StandardOrderParamsTypeDateRange `xml:"DateRange,omitempty"`
}

func NewStandardOrderParamsType() *StandardOrderParamsType {
	return new(StandardOrderParamsType)
}

func (me *StandardOrderParamsType) SetDateRange(value *StandardOrderParamsTypeDateRange) {
	me.DateRange = value
}

func (me *StandardOrderParamsType) AddDateRange() *StandardOrderParamsTypeDateRange {
	me.DateRange = new(StandardOrderParamsTypeDateRange)
	return me.DateRange
}
