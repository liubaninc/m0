package types

const (
	EventTypeDeploy  = "deploy"
	EventTypeUpgrade = "upgrade"
	EventTypeInvoke  = "invoke"

	EventTypeFreeze   = "freeze"
	EventTypeUnFreeze = "unfreeze"
	EventTypeUnDeploy = "undeploy"

	AttributeKeyName   = "name"
	AttributeKeyMethod = "method"
	AttributeKeyArg    = "arg"

	AttributeValueCategory = ModuleName
)
