package types

const (
	// ModuleName defines the module name
	ModuleName = "sender"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)

var (
	// SyncsKey defines the key to store all the synchronizations
	SyncsKey = []byte{0x01}
)

func NewSyncKey(id string) []byte {
	return append(SyncsKey, id...)
}
