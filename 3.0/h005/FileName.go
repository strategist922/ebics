// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// Attribute
type FileName struct {
	Value *FileNameStringType `xml:"fileName,attr"`
}

func NewFileName() *FileName {
	return new(FileName)
}

func (me *FileName) SetValue(value *FileNameStringType) {
	me.Value = value
}
