package types

const (
	EventTypeDeployContract = "Deploy"
	EventTypeUpgradeContract = "Upgrade"
	EventTypeInvokeContract = "Invoke"

	AttributeKeyName = "name"
	AttributeKeyMethod    = "method"
	AttributeKeyArg    = "arg"


	AttributeValueCategory = ModuleName
)