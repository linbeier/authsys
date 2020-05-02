package fileauthservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	// TODO: Define logic for when you would like to initalize a new genesis
	//这里的owner是不对的，先试看看这是干嘛用的，是脱机保存吗？
	//iter := k.GetAccountIterator(ctx)
	for _, record := range data.FileAuthRecords {
		if k.Findaccount(ctx, record.Owner) {
			k.AddFileauth(ctx, record.Owner, record)
		} else {
			recordslice := []Fileauth{record}
			k.SetFileauth(ctx, record.Owner, recordslice)
		}
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	// TODO: Define logic for exporting state
	var records []Fileauth
	iter := k.GetAccountIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		record := k.GetFileauth(ctx, iter.Key())
		for i := 0; i < len(record); i++ {
			records = append(records, record[i])
		}

	}
	return GenesisState{FileAuthRecords: records}
}
