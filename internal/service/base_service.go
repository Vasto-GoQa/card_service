package service

import (
	pb "card_service/generated/proto"
	"card_service/internal/repository"
)

type CardServiceServer struct {
	pb.UnimplementedCardServiceServer
	userRepo        *repository.UserRepository
	cardRepo        *repository.CardRepository
	transactionRepo *repository.TransactionRepository
}

func NewCardServiceServer(
	userRepo *repository.UserRepository,
	cardRepo *repository.CardRepository,
	transactionRepo *repository.TransactionRepository,
) *CardServiceServer {
	return &CardServiceServer{
		userRepo:        userRepo,
		cardRepo:        cardRepo,
		transactionRepo: transactionRepo,
	}
}
