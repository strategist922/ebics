// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

// ComplexType
type DateRangeType struct {
	Start *DateType `xml:"Start"`
	End   *DateType `xml:"End"`
}

func NewDateRangeType() *DateRangeType {
	return new(DateRangeType)
}

func (me *DateRangeType) SetStart(value *DateType) {
	me.Start = value
}

func (me *DateRangeType) SetEnd(value *DateType) {
	me.End = value
}
