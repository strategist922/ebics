// Generated with goxc v - rev
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex element
type AttributedAccountTypeAccountHolder struct {
	Value       AccountHolderType     `xml:",chardata"`
	Role        AccountHolderRoleType `xml:"Role,attr"`
	Description w3c.NormalizedString  `xml:"Description,attr"`
}