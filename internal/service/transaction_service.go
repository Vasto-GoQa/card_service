package service

import (
	pb "card_service/generated/proto"
	"card_service/internal/models"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CardServiceServer) CreateTransaction(_ context.Context, req *pb.CreateTransactionRequest) (*pb.Transaction, error) {
	transaction := &models.Transaction{
		FromCardID: int(req.FromCardId),
		ToCardID:   int(req.ToCardId),
		Amount:     req.Amount,
	}

	if req.Amount <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Amount must be greater than zero")
	}

	if req.FromCardId == req.ToCardId {
		return nil, status.Errorf(codes.InvalidArgument, "FromCardId and ToCardId cannot be the same")
	}

	createdTransaction, err := s.transactionRepo.Create(transaction)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create user: %v", err)
	}

	return s.modelTransactionToProto(createdTransaction), nil
}

func (s *CardServiceServer) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	transaction, err := s.transactionRepo.GetById(int(req.FromCardId), int(req.ToCardId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get transaction: %v", err)
	}

	if len(transaction) == 0 {
		return nil, status.Errorf(codes.NotFound, "No transactions found for the given cardIDs")
	}

	return &pb.GetTransactionResponse{
		Transactions: s.modelTransactionsToProto(transaction),
	}, nil
}
