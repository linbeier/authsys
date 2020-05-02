package fileauthservice

import (
	"github.com/linbeier/authsys/x/fileauthservice/internal/keeper"
	"github.com/linbeier/authsys/x/fileauthservice/internal/types"
)

const (
	// TODO: define constants that you would like exposed from the internal package

	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	RecordStoreKey    = types.RecordStoreKey
	DefaultParamspace = types.DefaultParamspace
	// QueryParams       = types.QueryParams
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	NewMsgSetFileAuth   = types.NewMsgSetFileAuth
	NewMsgTransFileAuth = types.NewMsgTransFileAuth
	NewMsgDelFileAuth   = types.NewMsgDelFileAuth
	// TODO: Fill out function aliases

	// variable aliases
	ModuleCdc = types.ModuleCdc
	// TODO: Fill out variable aliases
)

type (
	Keeper            = keeper.Keeper
	GenesisState      = types.GenesisState
	Params            = types.Params
	MsgSetFileAuth    = types.MsgSetFileAuth
	MsgTransFileAuth  = types.MsgTransFileAuth
	MsgDelFileAuth    = types.MsgDelFileAuth
	QueryResAuth      = types.QueryResAuth
	QueryResFileNames = types.QueryResFileNames
	QueryResAccounts  = types.QueryResAccounts
	Fileauth          = types.Fileauth
	Filerecord        = types.Filerecord
	// TODO: Fill out module types
)
