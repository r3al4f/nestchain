package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/r3al4f/nestchain/x/nestchain/types"
)

func (k msgServer) UpdateLoyaltyPoints(ctx context.Context, msg *types.MsgUpdateLoyaltyPoints) (*types.MsgUpdateLoyaltyPointsResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgUpdateLoyaltyPointsResponse{}, nil
}
