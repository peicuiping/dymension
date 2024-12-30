package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

const TypeMsgCreateRollapp = "create_rollapp"

var (
	_ sdk.Msg            = &MsgCreateRollapp{}
	_ legacytx.LegacyMsg = &MsgCreateRollapp{}
)

func NewMsgCreateRollapp(
	creator,
	rollappId,
	initSequencer string,
	minSequencerBond sdk.Coin,
	alias string,
	vmType Rollapp_VMType,
	metadata *RollappMetadata,
	genesisInfo *GenesisInfo,
) *MsgCreateRollapp {
	return &MsgCreateRollapp{
		Creator:          creator,
		RollappId:        rollappId,
		InitialSequencer: initSequencer,
		MinSequencerBond: minSequencerBond,
		Alias:            alias,
		VmType:           vmType,
		Metadata:         metadata,
		GenesisInfo:      genesisInfo,
	}
}

func (msg *MsgCreateRollapp) Route() string {
	return RouterKey
}

func (msg *MsgCreateRollapp) Type() string {
	return TypeMsgCreateRollapp
}

func (msg *MsgCreateRollapp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRollapp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRollapp) GetRollapp() Rollapp {
	genInfo := GenesisInfo{}
	if msg.GenesisInfo != nil {
		genInfo = *msg.GenesisInfo
		// hotfix: if supply is zero, override the denom metadata with empty
		if genInfo.InitialSupply.IsZero() {
			genInfo.NativeDenom = DenomMetadata{}
		}
	}
	return NewRollapp(
		msg.Creator,
		msg.RollappId,
		msg.InitialSequencer,
		msg.MinSequencerBond,
		msg.VmType,
		msg.Metadata,
		genInfo,
	)
}

func (msg *MsgCreateRollapp) ValidateBasic() error {
	if len(msg.Alias) == 0 {
		return ErrInvalidAlias
	}

	rollapp := msg.GetRollapp()
	if err := rollapp.ValidateBasic(); err != nil {
		return err
	}

	return nil
}
