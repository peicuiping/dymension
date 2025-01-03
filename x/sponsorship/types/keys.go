package types

import (
	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Module name and store keys.
const (
	// ModuleName defines the module name
	ModuleName = "sponsorship"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	RouterKey = ModuleName
)

const (
	ParamsByte                  = iota // Module params: Params
	DistributionByte                   // Current distribution: Distribution
	DelegatorValidatorPowerByte        // Delegator voting power by the validator: math.Int
	VoteByte                           // User's vote: Vote
)

func ParamsPrefix() collections.Prefix {
	return collections.NewPrefix(ParamsByte)
}

func DistributionPrefix() collections.Prefix {
	return collections.NewPrefix(DistributionByte)
}

func DelegatorValidatorPrefix() collections.Prefix {
	return collections.NewPrefix(DelegatorValidatorPowerByte)
}

func VotePrefix() collections.Prefix {
	return collections.NewPrefix(VoteByte)
}

func ParamsKey() []byte {
	return []byte{ParamsByte}
}

func DistributionKey() []byte {
	return []byte{DistributionByte}
}

func DelegatorValidatorPowerKey(voterAddr sdk.AccAddress, valAddr sdk.ValAddress) []byte {
	key := make([]byte, 0, 1+len(voterAddr)+len(valAddr))
	key = append(key, DelegatorValidatorPowerByte)
	key = append(key, voterAddr.Bytes()...)
	key = append(key, valAddr.Bytes()...)
	return key
}

func AllDelegatorValidatorPowersKey(voterAddr sdk.AccAddress) []byte {
	key := make([]byte, 0, 1+len(voterAddr))
	key = append(key, DelegatorValidatorPowerByte)
	key = append(key, voterAddr.Bytes()...)
	return key
}

func VoteKey(voterAddr sdk.AccAddress) []byte {
	return append([]byte{VoteByte}, voterAddr.Bytes()...)
}
