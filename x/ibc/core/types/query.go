package types

import (
	"github.com/gogo/protobuf/grpc"

	client "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/02-client"
	clienttypes "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/02-client/types"
	connection "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/03-connection"
	connectiontypes "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/03-connection/types"
	channel "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/04-channel"
	channeltypes "https://github.com/ArjavJP/Cosmos-sdk/x/ibc/core/04-channel/types"
)

// QueryServer defines the IBC interfaces that the gRPC query server must implement
type QueryServer interface {
	clienttypes.QueryServer
	connectiontypes.QueryServer
	channeltypes.QueryServer
}

// RegisterQueryService registers each individual IBC submodule query service
func RegisterQueryService(server grpc.Server, queryService QueryServer) {
	client.RegisterQueryService(server, queryService)
	connection.RegisterQueryService(server, queryService)
	channel.RegisterQueryService(server, queryService)
}
