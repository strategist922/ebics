// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type Parameter struct {
	Name  *w3c.Token      `xml:"Name"`
	Value *ParameterValue `xml:"Value"`
}

func NewParameter() *Parameter {
	return new(Parameter)
}

func (me *Parameter) SetName(value *w3c.Token) {
	me.Name = value
}

func (me *Parameter) SetValue(value *ParameterValue) {
	me.Value = value
}
