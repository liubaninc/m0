package types

// bank module event types
const (
	EventTypeSend    = "send"
	EventTypeIssue   = "issue"
	EventTypeDestory = "destory"

	AttributeKeyRecipient     = "recipient"
	AttributeKeySender        = "sender"
	AttributeKeyAmountChanged = "changed"
	AttributeKeyDenom         = "denom"

	AttributeValueCategory = ModuleName
)
