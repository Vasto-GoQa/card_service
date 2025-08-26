package card_service

import (
	cardService "card_service/generated/proto"
	"card_service/test/utils/config"
	"card_service/test/utils/logger"
	"card_service/test/utils/services/grpc"
)

type Client struct {
	cardService.CardServiceClient
	logger *logger.CtxLogger
}

func NewClient(logger logger.Service, conf config.Config) (*Client, error) {
	conn, err := grpc.GetGrpcClient(conf.Cardservice)

	if err != nil {
		return &Client{}, err
	}

	return &Client{
		logger:            logger.NewPrefix("CARD_SERVICE.GRPC.CLIENT"),
		CardServiceClient: cardService.NewCardServiceClient(conn),
	}, nil
}
