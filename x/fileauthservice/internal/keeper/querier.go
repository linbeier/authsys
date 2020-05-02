package keeper

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/linbeier/authsys/x/fileauthservice/internal/types"
)

const (
	QueryAuth        = "authority"
	QueryFile        = "files"
	QueryAccount     = "accounts"
	QueryRecords     = "records"
	QueryAll         = "all"
	QueryTraceRecord = "trace"
)

// NewQuerier creates a new querier for fileauthservice clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// case types.QueryParams:
		// 	return queryParams(ctx, k)
		// TODO: Put the modules query routes
		case QueryAccount:
			return queryAccount(ctx, req, k)
		case QueryFile:
			return queryFile(ctx, path[1:], req, k)
		case QueryAuth:
			return queryAuth(ctx, path[1:], req, k)
		case QueryAll:
			return queryAll(ctx, req, k)
		case QueryRecords:
			return queryRecord(ctx, path[1:], req, k)
		case QueryTraceRecord:
			return queryTraceRecord(ctx, path[1:], req, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("unknown fileauthservice query endpoint, path[0]: %s, path[1]: %s", path[0], path[1]))
		}
	}
}

// func queryParams(ctx sdk.Context, k Keeper) ([]byte, error) {
// 	params := k.GetParams(ctx)

// 	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
// 	if err != nil {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
// 	}

// 	return res, nil
// }

// TODO: Add the modules query functions
// They will be similar to the above one: queryParams()

func queryAccount(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var accountlist types.QueryResAccounts

	iterator := k.GetAccountIterator(ctx)

	var str string
	for ; iterator.Valid(); iterator.Next() {
		str = string(iterator.Key())
		fmt.Printf("iterator.key: %s\n", str)
		accountlist = append(accountlist, str)
	}

	res, err := codec.MarshalJSONIndent(k.cdc, accountlist)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

//querytype/account->
func queryFile(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {

	fp, err0 := os.OpenFile(fmt.Sprintf("/home/lin/go/src/github.com/linbeier/authsys/data.json"), os.O_CREATE|os.O_RDWR, 0755)
	if err0 != nil {
		fmt.Printf("%s", err0)
	}
	defer fp.Close()

	var keymap map[string][]byte
	keymap = make(map[string][]byte)
	keymapbyte, err0 := ioutil.ReadAll(fp)
	fmt.Printf("map byte: %s\n", keymapbyte)
	if err0 != nil {
		fmt.Printf("error with read\n")
	}
	if !bytes.Equal(nil, keymapbyte) {
		codec.Cdc.MustUnmarshalJSON(keymapbyte, &keymap)
	}
	var acc sdk.AccAddress
	acc = keymap[path[0]]

	// acc, err0 := sdk.AccAddressFromBech32(path[0])
	// if err0 != nil {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, err0.Error())
	// }

	fmt.Printf("acc convert: %s\n", acc.Bytes())

	filenames := k.Getfilenames(ctx, acc)

	fmt.Printf("fetch filename: %s\n", filenames)

	res, err := codec.MarshalJSONIndent(k.cdc, filenames)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

//querytype/account/filename/filehash->
func queryAuth(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {

	fp, err0 := os.OpenFile(fmt.Sprintf("/home/lin/go/src/github.com/linbeier/authsys/data.json"), os.O_CREATE|os.O_RDWR, 0755)
	if err0 != nil {
		fmt.Printf("%s", err0)
	}
	defer fp.Close()

	var keymap map[string][]byte
	keymap = make(map[string][]byte)
	keymapbyte, err0 := ioutil.ReadAll(fp)
	fmt.Printf("map byte: %s\n", keymapbyte)
	if err0 != nil {
		fmt.Printf("error with read\n")
	}
	if !bytes.Equal(nil, keymapbyte) {
		codec.Cdc.MustUnmarshalJSON(keymapbyte, &keymap)
	}
	var acc sdk.AccAddress
	acc = keymap[path[0]]
	// if err0 != nil {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, err0.Error())
	// }

	auth := k.Getauth(ctx, acc, path[1], path[2])

	res, err := codec.MarshalJSONIndent(k.cdc, auth)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

//time limited
func queryAll(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {

	return nil, nil

}

func queryRecord(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	fp, err0 := os.OpenFile(fmt.Sprintf("/home/lin/go/src/github.com/linbeier/authsys/data.json"), os.O_CREATE|os.O_RDWR, 0755)
	if err0 != nil {
		fmt.Printf("%s", err0)
	}
	defer fp.Close()

	var keymap map[string][]byte
	keymap = make(map[string][]byte)
	keymapbyte, err0 := ioutil.ReadAll(fp)
	fmt.Printf("map byte: %s\n", keymapbyte)
	if err0 != nil {
		fmt.Printf("error with read\n")
	}
	if !bytes.Equal(nil, keymapbyte) {
		codec.Cdc.MustUnmarshalJSON(keymapbyte, &keymap)
	}
	var acc sdk.AccAddress
	acc = keymap[path[0]]

	record := k.GetFilerecord(ctx, acc)

	res, err := codec.MarshalJSONIndent(k.cdc, record)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryTraceRecord(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	fp, err0 := os.OpenFile(fmt.Sprintf("/home/lin/go/src/github.com/linbeier/authsys/data.json"), os.O_CREATE|os.O_RDWR, 0755)
	if err0 != nil {
		fmt.Printf("%s", err0)
	}
	defer fp.Close()

	var keymap map[string][]byte
	keymap = make(map[string][]byte)
	keymapbyte, err0 := ioutil.ReadAll(fp)
	fmt.Printf("map byte: %s\n", keymapbyte)
	if err0 != nil {
		fmt.Printf("error with read\n")
	}
	if !bytes.Equal(nil, keymapbyte) {
		codec.Cdc.MustUnmarshalJSON(keymapbyte, &keymap)
	}
	var acc sdk.AccAddress
	acc = keymap[path[0]]

	records := k.GetFilerecord(ctx, acc)

	for i := 0; i < len(records); i++ {
		if records[i].Name == path[1] {
			record := records[i]
			res, err := codec.MarshalJSONIndent(k.cdc, record)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
			}

			return res, nil
		}
	}
	return nil, errors.New("no record found")
}
