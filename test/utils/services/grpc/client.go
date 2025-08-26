package grpc

import (
	"card_service/test/utils/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetGrpcClient(grpcService config.GrpcService) (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%d", grpcService.Host, grpcService.Port)

	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
