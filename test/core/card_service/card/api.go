package card_service

import (
	cardService "card_service/generated/proto"
	"context"
)

func (c *Client) getCard(
	ctx context.Context,
	request *cardService.GetCardRequest,
) (*cardService.Card, error) {
	c.logger.Sugar().Infof("GetCardRequest: %+v", request)

	res, err := c.CardServiceClient.GetCard(ctx, request)
	if err != nil {
		c.logger.Sugar().Errorf("GetCard failed: %v", err)
		return nil, err
	}

	c.logger.Sugar().Infof("CardResponse: %+v", res)
	return res, nil
}

func (c *Client) createCard(
	ctx context.Context,
	request *cardService.CreateCardRequest,
) (*cardService.Card, error) {
	c.logger.Sugar().Infof("CreateCardRequest: %+v", request)

	res, err := c.CardServiceClient.CreateCard(ctx, request)
	if err != nil {
		c.logger.Sugar().Errorf("CreateCard failed: %v", err)
		return nil, err
	}

	c.logger.Sugar().Infof("CreateCardResponse: %+v", res)
	return res, nil
}

func (c *Client) deleteCard(
	ctx context.Context,
	request *cardService.DeleteCardRequest,
) (*cardService.DeleteCardResponse, error) {
	c.logger.Sugar().Infof("DeleteCardRequest: %+v", request)

	res, err := c.CardServiceClient.DeleteCard(ctx, request)
	if err != nil {
		c.logger.Sugar().Errorf("DeleteCard failed: %v", err)
		return nil, err
	}

	c.logger.Sugar().Infof("DeleteCardResponse: %+v", res)
	return res, nil
}
