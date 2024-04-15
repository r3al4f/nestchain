package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/r3al4f/nestchain/x/nestchain/types"
)

func (k msgServer) CreateBooking(ctx context.Context, msg *types.MsgCreateBooking) (*types.MsgCreateBookingResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgCreateBookingResponse{}, nil
}
