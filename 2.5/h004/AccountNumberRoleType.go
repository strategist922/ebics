// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

type AccountNumberRoleType w3c.Token

const (
	AccountNumberRoleType_ORIGINATOR AccountNumberRoleType = "Originator"
	AccountNumberRoleType_RECIPIENT  AccountNumberRoleType = "Recipient"
	AccountNumberRoleType_CHARGES    AccountNumberRoleType = "Charges"
	AccountNumberRoleType_OTHER      AccountNumberRoleType = "Other"
)
