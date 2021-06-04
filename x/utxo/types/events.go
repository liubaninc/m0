package types

// bank module event types
const (
	EventTypeTransfer = "transfer"

	AttributeKeyCreator       = "creator"
	AttributeKeyRecipient     = "recipient"
	AttributeKeySender        = "sender"
	AttributeKeyAmountChanged = "changed"
	AttributeKeyDenom         = "denom"

	AttributeValueCategory = ModuleName
)
