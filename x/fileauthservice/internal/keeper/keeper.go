package keeper

import (
	"fmt"
	"strings"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/linbeier/authsys/x/fileauthservice/internal/types"
)

//介于现在时间比较短，先把底层存储按账户-权限集数组的方式来存储。之后希望能够重构为账户-文件map的形式，一可以方便查询，二可以避免重复造轮子

// Keeper of the fileauthservice store
type Keeper struct {
	CoinKeeper bank.Keeper
	storeKey   sdk.StoreKey
	recordKey  sdk.StoreKey
	cdc        *codec.Codec
	//paramspace types.ParamSubspace
}

// NewKeeper creates a fileauthservice keeper
func NewKeeper(cdc *codec.Codec, storekey sdk.StoreKey, recordkey sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey:  storekey,
		recordKey: recordkey,
		cdc:       cdc,
		// paramspace: paramspace.WithKeyTable(types.ParamKeyTable()),
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// // Get returns the pubkey from the adddress-pubkey relation
// func (k Keeper) Get(ctx sdk.Context, key string) (/* TODO: Fill out this type */, error) {
// 	store := ctx.KVStore(k.storeKey)
// 	var item /* TODO: Fill out this type */
// 	byteKey := []byte(key)
// 	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &item)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return item, nil
// }

// func (k Keeper) set(ctx sdk.Context, key string, value /* TODO: fill out this type */ ) {
// 	store := ctx.KVStore(k.storeKey)
// 	bz := k.cdc.MustMarshalBinaryLengthPrefixed(value)
// 	store.Set([]byte(key), bz)
// }

func (k Keeper) delete(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(key))
}

//key:account, value: fileauth struct
func (k Keeper) SetFileauth(ctx sdk.Context, key sdk.AccAddress, file []types.Fileauth) {
	if file[0].Name == "" || file[0].Hash == "" || file[0].Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(key), k.cdc.MustMarshalBinaryBare(file))
	fmt.Printf("set store: %s\n", string(k.cdc.MustMarshalBinaryBare(file)))
}

func (k Keeper) GetFileauth(ctx sdk.Context, key sdk.AccAddress) []types.Fileauth {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(key)) {
		return nil
	}
	bz := store.Get([]byte(key))
	fmt.Printf("fetched bytes: %s\n", bz)
	var file []types.Fileauth
	k.cdc.MustUnmarshalBinaryBare(bz, &file)
	fmt.Printf("get unmarshal: %s %s\n", file[0].Name, file[0].Owner)
	return file
}

//所有底层存储都忘了unmarshal了！！！！
func (k Keeper) AddFileauth(ctx sdk.Context, key sdk.AccAddress, file types.Fileauth) {
	if file.Name == "" || file.Hash == "" || file.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)

	var fileauthSlice []types.Fileauth
	var target = -1

	fileauthSlice = k.GetFileauth(ctx, key)
	for i := 0; i < len(fileauthSlice); i++ {
		if file.Hash == fileauthSlice[i].Hash {
			target = i
			break
		}
	}

	if target == -1 {
		fileauthSlice = append(fileauthSlice, file)
	} else {
		authIni := fileauthSlice[target].Auth
		auth := file.Auth
		for _, ch := range auth {

			if strings.Contains(authIni, string(ch)) {
				continue
			} else {
				authIni += string(ch)
			}
		}
		fileauthSlice[target].Auth = authIni
	}
	store.Set([]byte(key), k.cdc.MustMarshalBinaryBare(fileauthSlice))
	return
}

func (k Keeper) DelFileauth(ctx sdk.Context, key sdk.AccAddress, file types.Fileauth) {
	if file.Name == "" || file.Hash == "" || file.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)

	var fileauthSlice []types.Fileauth
	var target = -1
	fileauthSlice = k.GetFileauth(ctx, key)
	for i := 0; i < len(fileauthSlice); i++ {
		if file.Hash == fileauthSlice[i].Hash {
			target = i
			break
		}
	}

	if target == -1 {
		return
	} else {
		authIni := fileauthSlice[target].Auth
		auth := file.Auth
		for _, ch := range auth {

			if strings.Contains(authIni, string(ch)) {
				authIni = strings.Replace(authIni, string(ch), "", -1)
			} else {
				continue
			}
		}
		fileauthSlice[target].Auth = authIni
	}
	store.Set([]byte(key), k.cdc.MustMarshalBinaryBare(fileauthSlice))
}

/* interface of KVstore
interface {
    Delete(key []byte)
    Get(key []byte) []byte
    Has(key []byte) bool
    Iterator(start []byte, end []byte) Iterator
    ReverseIterator(start []byte, end []byte) Iterator
    Set(key []byte, value []byte)
    Store
}

type Fileauth struct {
	Name  string		 `json:"name"`
	Hash  string         `json:"hash"`
	Owner sdk.AccAddress `json:"owner"`
	Auth  string		 `json:"auth"`
}
*/
func (k Keeper) Findaccount(ctx sdk.Context, key sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(key))

}

func (k Keeper) Findfile(ctx sdk.Context, key sdk.AccAddress, name string, hash string) bool {
	var fileauthSlice []types.Fileauth
	fileauthSlice = k.GetFileauth(ctx, key)

	for i := 0; i < len(fileauthSlice); i++ {
		if hash == fileauthSlice[i].Hash && name == fileauthSlice[i].Name {
			return true
		}
	}

	return false
}

func (k Keeper) Getfilenames(ctx sdk.Context, key sdk.AccAddress) []string {
	var fileauthSlice []types.Fileauth
	fileauthSlice = k.GetFileauth(ctx, key)

	var filenames types.QueryResFileNames
	for i := 0; i < len(fileauthSlice); i++ {
		filenames = append(filenames, fileauthSlice[i].Name)
	}

	return filenames
}

func (k Keeper) Getauth(ctx sdk.Context, key sdk.AccAddress, name string, hash string) string {
	var fileauthSlice []types.Fileauth
	fileauthSlice = k.GetFileauth(ctx, key)

	for i := 0; i < len(fileauthSlice); i++ {
		fmt.Printf("name: %s\n hash: %s\n", fileauthSlice[i].Name, fileauthSlice[i].Hash)
		if (hash == fileauthSlice[i].Hash) && (name == fileauthSlice[i].Name) {
			return fileauthSlice[i].Auth
		}
	}

	return ""
}

func (k Keeper) GetOwner(ctx sdk.Context, key sdk.AccAddress, name string, hash string) sdk.AccAddress {
	var fileauthSlice []types.Fileauth
	fileauthSlice = k.GetFileauth(ctx, key)

	for i := 0; i < len(fileauthSlice); i++ {
		if hash == fileauthSlice[i].Hash && name == fileauthSlice[i].Name {
			return fileauthSlice[i].Owner
		}
	}

	return nil
}

// func (k Keeper) GetPrice(ctx sdk.Context, key sdk.AccAddress, name string, hash string) sdk.Coins {
// 	var fileauthSlice []Fileauth
// 	fileauthSlice = k.GetFileauth(ctx, key)

// 	for i := 0; i < len(fileauthSlice); i++ {
// 		if hash == fileauthSlice[i].Hash && name == fileauthSlice[i].Name {
// 			return fileauthSlice[i].Price
// 		}
// 	}
// 	return nil
// }

func (k Keeper) Judgeauth(ctx sdk.Context, key sdk.AccAddress, name string, hash string, auth string) bool {
	var fileauthSlice []types.Fileauth
	fileauthSlice = k.GetFileauth(ctx, key)

	authbyte := []byte(auth)
	for i := 0; i < len(fileauthSlice); i++ {
		if hash == fileauthSlice[i].Hash && name == fileauthSlice[i].Name {
			for j := 0; j < len(authbyte); j++ {
				if !strings.Contains(fileauthSlice[i].Auth, string(authbyte[j])) {
					return false
				}
			}
		}
	}
	return true
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetAccountIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	if store == nil {
		fmt.Println("GetAccountIterator error!")
	}
	return sdk.KVStorePrefixIterator(store, []byte{})
}

//************针对数据转移的存储操作，操作对像为reocrdstrore*******************//

func (k Keeper) SetFilerecord(ctx sdk.Context, key sdk.AccAddress, record []types.Filerecord) {
	store := ctx.KVStore(k.recordKey)
	if store.Has([]byte(key)) {
		k.AddFilerecord(ctx, key, record)
	} else {
		store.Set([]byte(key), k.cdc.MustMarshalBinaryBare(record))
		fmt.Printf("set recordstore: %s\n", string(k.cdc.MustMarshalBinaryBare(record)))
	}

}

func (k Keeper) AddFilerecord(ctx sdk.Context, key sdk.AccAddress, record []types.Filerecord) {
	store := ctx.KVStore(k.recordKey)
	records := k.GetFilerecord(ctx, key)
	records = append(records, record[0])
	store.Set([]byte(key), k.cdc.MustMarshalBinaryBare(records))
}

func (k Keeper) GetFilerecord(ctx sdk.Context, key sdk.AccAddress) []types.Filerecord {
	store := ctx.KVStore(k.recordKey)
	if !store.Has([]byte(key)) {
		return nil
	}
	bz := store.Get([]byte(key))
	fmt.Printf("fetched bytes: %s\n", bz)
	var file []types.Filerecord
	k.cdc.MustUnmarshalBinaryBare(bz, &file)
	fmt.Printf("get unmarshal: %s %s\n", file[0].Name, file[0].From)
	return file
}
