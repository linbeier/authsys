package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Fill out some custom errors for the module
// You can see how they are constructed below:
// var (
//	ErrInvalid = sdkerrors.Register(ModuleName, 1, "custom error message")
// )

var (
	ErrFileDoesNotExist = sdkerrors.Register(ModuleName, 1, "File does not exist")
	ErrNoTransferRight  = sdkerrors.Register(ModuleName, 2, "The Account cannot transfer this file's rights")
	//TO DO

)
