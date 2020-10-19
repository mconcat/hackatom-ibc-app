package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	abci "github.com/tendermint/tendermint/abci/types"
	cryptoencoding "github.com/tendermint/tendermint/crypto/encoding"
)

func MustUnmarshalValidatorUpdate(cdc codec.BinaryMarshaler, value []byte) abci.ValidatorUpdate {
	validator, err := UnmarshalValidatorUpdate(cdc, value)
	if err != nil {
		panic(err)
	}
	return validator
}

func UnmarshalValidatorUpdate(cdc codec.BinaryMarshaler, value []byte) (v abci.ValidatorUpdate, err error) {
	err = cdc.UnmarshalBinaryBare(value, &v)
	return v, err
}

func ToABCIValidator(vu abci.ValidatorUpdate) abci.Validator {
	pubkey, err := cryptoencoding.PubKeyFromProto(vu.PubKey)
	if err != nil {
		panic(err) // TODO
	}
	return abci.Validator{
		Address: pubkey.Address(),
		Power:   vu.Power,
	}
}
