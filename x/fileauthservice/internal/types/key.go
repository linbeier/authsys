package types

const (
	// ModuleName is the name of the module
	ModuleName = "fileauthservice"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	//RecordStoreKey to be used in store tracing message
	RecordStoreKey = "recordservice"

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName
)
