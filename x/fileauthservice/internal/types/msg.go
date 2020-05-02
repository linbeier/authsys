package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`
/*
verify interface at compile time
var _ sdk.Msg = &Msg<Action>{}

Msg<Action> - struct for unjailing jailed validator
type Msg<Action> struct {
	ValidatorAddr sdk.ValAddress `json:"address" yaml:"address"` // address of the validator operator
}

NewMsg<Action> creates a new Msg<Action> instance
func NewMsg<Action>(validatorAddr sdk.ValAddress) Msg<Action> {
	return Msg<Action>{
		ValidatorAddr: validatorAddr,
	}
}

const <action>Const = "<action>"

// nolint
func (msg Msg<Action>) Route() string { return RouterKey }
func (msg Msg<Action>) Type() string  { return <action>Const }
func (msg Msg<Action>) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ValidatorAddr)}
}

GetSignBytes gets the bytes for the message signer to sign on
func (msg Msg<Action>) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

ValidateBasic validity check for the AnteHandler
func (msg Msg<Action>) ValidateBasic() error {
	if msg.ValidatorAddr.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing validator address"
	}
	return nil
}
*/
type MsgSetFileAuth struct {
	Name   string         `json:"name"`
	Hash   string         `json:"hash"`
	Owner  sdk.AccAddress `json:"owner"`
	Origin sdk.AccAddress `json:"origin"`
	Auth   string         `json:"auth"`
}

func NewMsgSetFileAuth(name string, hash string, owner sdk.AccAddress, origin sdk.AccAddress, auth string) MsgSetFileAuth {
	return MsgSetFileAuth{
		Name:   name,
		Hash:   hash,
		Owner:  owner,
		Origin: origin,
		Auth:   auth,
	}
}

// Route should return the name of the module
func (msg MsgSetFileAuth) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetFileAuth) Type() string { return "set_fileauth" }

//The above functions are used by the SDK to route Msgs to the proper module for handling. They also add human readable names to database tags used for indexing.

//ValidateBasic is used to provide some basic stateless checks on the validity of the Msg. In this case, check that none of the attributes are empty.
func (msg MsgSetFileAuth) ValidateBasic() error {
	if msg.Owner.Empty() || msg.Origin.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Hash) == 0 || len(msg.Auth) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "name/hash/authority cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetFileAuth) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

//GetSigners defines whose signature is required on a Tx in order for it to be valid. In this case, for example, the MsgSetName requires that the Owner signs the transaction when trying to reset what the name points to.
func (msg MsgSetFileAuth) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

//This is a message that will transfer a file's authority to another account. Additionally, the owner will not lose the authority
//For auth, it is the auth to be added
type MsgTransFileAuth struct {
	Name      string         `json:"name"`
	Hash      string         `json:"hash"`
	Owner     sdk.AccAddress `json:"owner"`
	ToAccount sdk.AccAddress `json:"toaccount"`
	Origin    sdk.AccAddress `json:"origin"`
	Auth      string         `json:"auth"`
}

func NewMsgTransFileAuth(name string, hash string, owner sdk.AccAddress, toaccount sdk.AccAddress, origin sdk.AccAddress, auth string) MsgTransFileAuth {
	return MsgTransFileAuth{
		Name:      name,
		Hash:      hash,
		Owner:     owner,
		ToAccount: toaccount,
		Origin:    origin,
		Auth:      auth,
	}
}

// Route should return the name of the module
func (msg MsgTransFileAuth) Route() string { return RouterKey }

// Type should return the action
func (msg MsgTransFileAuth) Type() string { return "transfer fileauth" }

//The above functions are used by the SDK to route Msgs to the proper module for handling. They also add human readable names to database tags used for indexing.

//ValidateBasic is used to provide some basic stateless checks on the validity of the Msg. In this case, check that none of the attributes are empty.
func (msg MsgTransFileAuth) ValidateBasic() error {
	if msg.Owner.Empty() || msg.ToAccount.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Hash) == 0 || len(msg.Auth) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "name/hash/authority cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgTransFileAuth) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

//GetSigners defines whose signature is required on a Tx in order for it to be valid. In this case, for example, the MsgSetName requires that the Owner signs the transaction when trying to reset what the name points to.
func (msg MsgTransFileAuth) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

//设计标准是owner才可以进行删除操作
type MsgDelFileAuth struct {
	Name      string         `json:"name"`
	Hash      string         `json:"hash"`
	Owner     sdk.AccAddress `json:"owner"`
	ToAccount sdk.AccAddress `json:"toaccount"`
	Origin    sdk.AccAddress `json:"origin"`
	Auth      string         `json:"auth"`
}

func NewMsgDelFileAuth(name string, hash string, owner sdk.AccAddress, toaccount sdk.AccAddress, origin sdk.AccAddress, auth string) MsgDelFileAuth {
	return MsgDelFileAuth{
		Name:      name,
		Hash:      hash,
		Owner:     owner,
		ToAccount: toaccount,
		Origin:    origin,
		Auth:      auth,
	}
}

// Route should return the name of the module
func (msg MsgDelFileAuth) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDelFileAuth) Type() string { return "transfer fileauth" }

//The above functions are used by the SDK to route Msgs to the proper module for handling. They also add human readable names to database tags used for indexing.

//ValidateBasic is used to provide some basic stateless checks on the validity of the Msg. In this case, check that none of the attributes are empty.
func (msg MsgDelFileAuth) ValidateBasic() error {
	if msg.Owner.Empty() || msg.ToAccount.Empty() || msg.Origin.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Hash) == 0 || len(msg.Auth) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "name/hash/authority cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDelFileAuth) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

//GetSigners defines whose signature is required on a Tx in order for it to be valid. In this case, for example, the MsgSetName requires that the Owner signs the transaction when trying to reset what the name points to.
func (msg MsgDelFileAuth) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
