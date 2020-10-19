package types

import (
	"sort"

	abci "github.com/tendermint/tendermint/abci/types"
)

var _ sort.Interface = ValidatorUpdates{}

func (vu ValidatorUpdates) Len() int {
	return len(vu.ValidatorUpdates)
}

func (vu ValidatorUpdates) Less(i, j int) bool {
	return vu.ValidatorUpdates[i].PubKey.Compare(vu.ValidatorUpdates[j].PubKey) == -1
}

func (vu ValidatorUpdates) Swap(i, j int) {
	orig := vu.ValidatorUpdates[i]
	vu.ValidatorUpdates[i] = vu.ValidatorUpdates[j]
	vu.ValidatorUpdates[j] = orig
}

func NewValidatorUpdates(updates []*abci.ValidatorUpdate) ValidatorUpdates {
	return ValidatorUpdates{updates}
}
