package service

import (
	pb "card_service/generated/proto"
	"card_service/internal/models"
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CardServiceServer) CreateCard(ctx context.Context, req *pb.CreateCardRequest) (*pb.Card, error) {
	if req.UserId < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "UserId must be provided and greater than zero")
	}

	issueDate, err := time.Parse("2006-01-02", req.IssueDate)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid issue date format, expected YYYY-MM-DD")
	}

	expiryDate, err := time.Parse("2006-01-02", req.ExpiryDate)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid expiry date format, expected YYYY-MM-DD")
	}

	if !expiryDate.After(issueDate) {
		return nil, status.Errorf(codes.InvalidArgument, "Expiry date must be after issue date")
	}

	if len(req.CardNumber) < 16 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid card number: must be at least 16 characters")
	}

	if req.OperatorId < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "OperatorId must be provided and greater than zero")
	}

	if req.Balance < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Balance cannot be negative")
	}

	if req.Balance >= 1e14 {
		return nil, status.Errorf(codes.InvalidArgument, "Balance cannot be negative")
	}

	card := &models.Card{
		UserID:     int(req.UserId),
		CardNumber: req.CardNumber,
		OperatorID: int(req.OperatorId),
		IssueDate:  issueDate,
		ExpiryDate: expiryDate,
		Balance:    req.Balance,
	}

	createdCard, err := s.cardRepo.Create(card)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create card: %v", err)
	}

	operator, err := s.cardRepo.GetOperatorByID(createdCard.OperatorID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve operator: %v", err)
	}

	return s.modelCardToProto(createdCard, operator), nil
}

func (s *CardServiceServer) GetCard(ctx context.Context, req *pb.GetCardRequest) (*pb.Card, error) {
	if req.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "Id must be provided and greater than zero")
	}

	card, err := s.cardRepo.GetByID(int(req.Id))
	if err != nil {
		// если репозиторий возвращает "не найдено"
		if errors.Is(err, models.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "card with id %d not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "Failed to get card: %v", err)
	}

	operator, err := s.cardRepo.GetOperatorByID(card.OperatorID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "operator with id %d not found", card.OperatorID)
		}
		return nil, status.Errorf(codes.Internal, "Failed to retrieve operator: %v", err)
	}

	return s.modelCardToProto(card, operator), nil
}

func (s *CardServiceServer) DeleteCard(ctx context.Context, req *pb.DeleteCardRequest) (*pb.DeleteCardResponse, error) {
	if req.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "Id must be provided and greater than zero")
	}

	err := s.cardRepo.Delete(int(req.Id))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "card with id %d not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "Failed to delete card: %v", err)
	}

	return &pb.DeleteCardResponse{
		Success: true,
		Message: "Card deleted successfully",
	}, nil
}
