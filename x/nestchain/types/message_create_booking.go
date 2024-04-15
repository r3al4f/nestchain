package types

func NewMsgCreateBooking(creator string, hotelId string, customerId string, startDate string, endDate string, roomType string, price string) *MsgCreateBooking {
	return &MsgCreateBooking{
		Creator:    creator,
		HotelId:    hotelId,
		CustomerId: customerId,
		StartDate:  startDate,
		EndDate:    endDate,
		RoomType:   roomType,
		Price:      price,
	}
}
