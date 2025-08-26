package card_service

import (
	cardService "card_service/generated/proto"
	"context"
)

func (c *Client) getTransaction(
	ctx context.Context,
	request *cardService.GetTransactionRequest,
) (*cardService.GetTransactionResponse, error) {
	c.logger.Sugar().Infof("GetTransactionRequest: %+v", request)

	res, err := c.CardServiceClient.GetTransaction(ctx, request)
	if err != nil {
		c.logger.Sugar().Errorf("GetTransaction failed: %v", err)
		return nil, err
	}

	c.logger.Sugar().Infof("TransactionResponse: %+v", res)
	return res, nil
}

func (c *Client) createTransaction(
	ctx context.Context,
	request *cardService.CreateTransactionRequest,
) (*cardService.Transaction, error) {
	c.logger.Sugar().Infof("CreateTransactionRequest: %+v", request)

	res, err := c.CardServiceClient.CreateTransaction(ctx, request)
	if err != nil {
		c.logger.Sugar().Errorf("CreateTransaction failed: %v", err)
		return nil, err
	}

	c.logger.Sugar().Infof("CreateTransactionResponse: %+v", res)
	return res, nil
}
