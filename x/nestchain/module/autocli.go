package nestchain

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/r3al4f/nestchain/api/nestchain/nestchain"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateBooking",
					Use:            "create-booking [hotel-id] [customer-id] [start-date] [end-date] [room-type] [price]",
					Short:          "Send a createBooking tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "hotelId"}, {ProtoField: "customerId"}, {ProtoField: "startDate"}, {ProtoField: "endDate"}, {ProtoField: "roomType"}, {ProtoField: "price"}},
				},
				{
					RpcMethod:      "UpdateLoyaltyPoints",
					Use:            "update-loyalty-points [customer-id] [points]",
					Short:          "Send a updateLoyaltyPoints tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "customerId"}, {ProtoField: "points"}},
				},
				{
					RpcMethod:      "ProcessPayment",
					Use:            "process-payment [customer-id] [amount] [description]",
					Short:          "Send a processPayment tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "customerId"}, {ProtoField: "amount"}, {ProtoField: "description"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
