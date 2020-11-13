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
	SyncEntriesKey = []byte{0x01}
	
	// ValidatorSetsKey defines the key to store validator set at specific height
	ValidatorSetsKey = []byte{0x02}
)

func NewSyncEntryKey(id string) []byte {
	return append(SyncsKey, id...)
}

func NewValidatorSetKey(id string, height uint64) []byte {
	heightBz := sdk.Uint64ToBigEndian(uint64(height))
	prefixL := len(ValidatorSetsKey)
	idL := len(id)
	bz := make([]byte, prefixL+idL+8)

	// TODO: check id validity, avoid possible collision with different lenght of ids

	copy(bz[:prefixL], ValidatorSetKey)
	copy(bz[prefixL:prefixL+idL], []byte(id))
	copy(bz[prefixL+idL:], heightBz)

	return append(ValidatorSetDiffsKey, bz...)
}
