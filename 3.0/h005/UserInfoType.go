// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type UserInfoType struct {
	UserID     *UserInfoTypeUserID   `xml:"UserID"`
	Name       *NameType             `xml:"Name,omitempty"`
	Permission []*UserPermissionType `xml:"Permission"`

	Any []*w3c.Any
}
