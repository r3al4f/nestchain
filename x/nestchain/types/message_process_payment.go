package types

func NewMsgProcessPayment(creator string, customerId string, amount string, description string) *MsgProcessPayment {
	return &MsgProcessPayment{
		Creator:     creator,
		CustomerId:  customerId,
		Amount:      amount,
		Description: description,
	}
}
