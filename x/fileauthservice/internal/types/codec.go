package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// TODO: Register the modules msgs
	cdc.RegisterConcrete(MsgSetFileAuth{}, "fileauthservice/SetFileAuth", nil)
	cdc.RegisterConcrete(MsgTransFileAuth{}, "fileauthservice/TransFileAuth", nil)
	cdc.RegisterConcrete(MsgDelFileAuth{}, "fileauthservice/DelFileAuth", nil)
	cdc.RegisterConcrete(Fileauth{}, "fileauthservice/Fileauth", nil)
	cdc.RegisterConcrete([]Fileauth{}, "fileauthservice/[]Fileauth", nil)
	cdc.RegisterConcrete([]Filerecord{}, "fileauthservice/[]Filerecord", nil)
	cdc.RegisterConcrete(QueryResAuth{}, "fileauthservice/QueryResAuth", nil)
	cdc.RegisterConcrete(QueryResFileNames{}, "fileauthservice/QueryResFileNames", nil)
	cdc.RegisterConcrete(QueryResAccounts{}, "fileauthservice/QueryResAccounts", nil)

}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
