package types

func NewMsgUpdateLoyaltyPoints(creator string, customerId string, points string) *MsgUpdateLoyaltyPoints {
	return &MsgUpdateLoyaltyPoints{
		Creator:    creator,
		CustomerId: customerId,
		Points:     points,
	}
}
