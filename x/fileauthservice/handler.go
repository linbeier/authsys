package fileauthservice

import (
	"errors"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the fileauthservice type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// TODO: Define your msg cases
		//
		//Example:
		// case Msg<Action>:
		// 	return handleMsg<Action>(ctx, k, msg)
		case MsgSetFileAuth:
			return handleMsgSetFileAuth(ctx, k, msg)
		case MsgTransFileAuth:
			return handleMsgTransFileAuth(ctx, k, msg)
		case MsgDelFileAuth:
			return handleMsgDelFileAuth(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// // handle<Action> does x
// func handleMsg<Action>(ctx sdk.Context, k Keeper, msg Msg<Action>) (*sdk.Result, error) {
// 	err := k.<Action>(ctx, msg.ValidatorAddr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// TODO: Define your msg events
// 	ctx.EventManager().EmitEvent(
// 		sdk.NewEvent(
// 			sdk.EventTypeMessage,
// 			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
// 			sdk.NewAttribute(sdk.AttributeKeySender, msg.ValidatorAddr.String()),
// 		),
// 	)

// 	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
// }

func handleMsgSetFileAuth(ctx sdk.Context, k Keeper, msg MsgSetFileAuth) (*sdk.Result, error) {
	//if !msg.Owner.Equals(Keeper.GetOwner(ctx,msg.Name))
	//keeper.SetName(ctx, msg.Name, msg.Value) // If so, set the name to the value specified in the msg.
	//return &sdk.Result{}, nil                // return
	//if !msg.Owner.Equals(k.GetOwner(ctx, msg.Owner, msg.Name, msg.Hash)){
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	//}
	msgre := Fileauth{msg.Name, msg.Hash, msg.Owner, msg.Origin, msg.Auth}

	k.AddFileauth(ctx, msg.Owner, msgre)

	return &sdk.Result{}, nil
}

func handleMsgTransFileAuth(ctx sdk.Context, k Keeper, msg MsgTransFileAuth) (*sdk.Result, error) {
	//conditions judgement?

	if k.Findfile(ctx, msg.Owner, msg.Name, msg.Hash) {
		if k.Judgeauth(ctx, msg.Owner, msg.Name, msg.Hash, msg.Auth) {

			msgre := Fileauth{msg.Name, msg.Hash, msg.ToAccount, msg.Origin, msg.Auth}

			k.AddFileauth(ctx, msg.ToAccount, msgre)
			nowtime := time.Now()
			msgrecord := Filerecord{msg.Name, msg.Hash, msg.Owner, msg.Origin, nowtime}
			msgrecords := []Filerecord{msgrecord}

			k.SetFilerecord(ctx, msg.ToAccount, msgrecords)

			// rec := k.GetFilerecord(ctx, msg.ToAccount)
			// fmt.Printf("%s\n", rec[0].String())
			return &sdk.Result{}, nil
		}
	}
	err := errors.New("transfer error")
	return nil, err
}

func handleMsgDelFileAuth(ctx sdk.Context, k Keeper, msg MsgDelFileAuth) (*sdk.Result, error) {
	// if !msg.Owner.Equals(k.GetOwner(ctx, msg.Owner, msg.Name, msg.Hash)) {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	// }
	if k.Judgeauth(ctx, msg.ToAccount, msg.Name, msg.Hash, msg.Auth) {
		msgre := Fileauth{msg.Name, msg.Hash, msg.Owner, msg.Origin, msg.Auth}
		k.DelFileauth(ctx, msg.ToAccount, msgre)
		return &sdk.Result{}, nil
	}
	err := errors.New("delete error")
	return nil, err
}
