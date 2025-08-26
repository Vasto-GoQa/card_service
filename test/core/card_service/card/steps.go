package card_service

import (
	proto "card_service/generated/proto"
	"context"
	"fmt"

	"github.com/dailymotion/allure-go"
)

func (c *Client) GetCard(ctx context.Context, request *proto.GetCardRequest) (*proto.Card, error) {
	var (
		response *proto.Card
		err      error
	)

	allure.Step(allure.Description("Send CardService.GetCard request"), allure.Action(func() {
		response, err = c.getCard(ctx, request)
	}))

	if err != nil {
		return nil, fmt.Errorf("GetCard failed: %w", err)
	}
	return response, nil
}

func (c *Client) CreateCard(ctx context.Context, request *proto.CreateCardRequest) (*proto.Card, int32, error) {
	var (
		response *proto.Card
		err      error
	)

	allure.Step(allure.Description("Send CardService.CreateCard request"), allure.Action(func() {
		response, err = c.createCard(ctx, request)
	}))

	if err != nil {
		return nil, 0, fmt.Errorf("CreateCard failed: %w", err)
	}

	cardID := response.GetId()
	allure.Step(allure.Description(fmt.Sprintf("Created card with ID: %d", cardID)), allure.Action(func() {}))

	return response, cardID, nil
}

func (c *Client) DeleteCard(ctx context.Context, request *proto.DeleteCardRequest) (*proto.DeleteCardResponse, error) {
	var (
		response *proto.DeleteCardResponse
		err      error
	)

	allure.Step(allure.Description("Send CardService.DeleteCard request"), allure.Action(func() {
		response, err = c.deleteCard(ctx, request)
	}))

	if err != nil {
		return nil, fmt.Errorf("DeleteCard failed: %w", err)
	}
	return response, nil
}
