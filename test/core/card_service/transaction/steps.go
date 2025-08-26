package card_service

import (
	proto "card_service/generated/proto"
	"context"
	"fmt"

	"github.com/dailymotion/allure-go"
)

func (c *Client) GetTransaction(ctx context.Context, request *proto.GetTransactionRequest) (*proto.GetTransactionResponse, error) {
	var (
		response *proto.GetTransactionResponse
		err      error
	)

	allure.Step(allure.Description("Send CardService.GetTransaction request"), allure.Action(func() {
		response, err = c.getTransaction(ctx, request)
	}))

	if err != nil {
		return nil, fmt.Errorf("GetTransaction failed: %w", err)
	}
	return response, nil
}

func (c *Client) CreateTransaction(ctx context.Context, request *proto.CreateTransactionRequest) (*proto.Transaction, int32, error) {
	var (
		response *proto.Transaction
		err      error
	)

	allure.Step(allure.Description("Send CardService.CreateTransaction request"), allure.Action(func() {
		response, err = c.createTransaction(ctx, request)
	}))

	if err != nil {
		return nil, 0, fmt.Errorf("CreateTransaction failed: %w", err)
	}

	transactionID := response.GetId()
	allure.Step(allure.Description(fmt.Sprintf("Created transaction with ID: %d", transactionID)), allure.Action(func() {}))

	return response, transactionID, nil
}
