package service

import (
	pb "card_service/generated/proto"
	"card_service/internal/models"
	"card_service/internal/repository"
	"context"
	"database/sql"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type CardServiceServer struct {
	pb.UnimplementedCardServiceServer
	userRepo *repository.UserRepository
	cardRepo *repository.CardRepository
}

func NewCardServiceServer(userRepo *repository.UserRepository, cardRepo *repository.CardRepository) *CardServiceServer {
	return &CardServiceServer{
		userRepo: userRepo,
		cardRepo: cardRepo,
	}
}

// CreateUser Операции с пользователями
func (s *CardServiceServer) CreateUser(_ context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	if req.Phone != "" {
		user.Phone = sql.NullString{String: req.Phone, Valid: true}
	}

	if req.BirthDate != "" {
		birthDate, err := time.Parse("2006-01-02", req.BirthDate)
		if err != nil {
			return &pb.UserResponse{Error: "Invalid birth date format"}, nil
		}
		user.BirthDate = sql.NullTime{Time: birthDate, Valid: true}
	}

	createdUser, err := s.userRepo.Create(user)
	if err != nil {
		return &pb.UserResponse{Error: err.Error()}, nil
	}

	return &pb.UserResponse{
		User: s.modelUserToProto(createdUser),
	}, nil
}

func (s *CardServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, err := s.userRepo.GetByID(int(req.Id))
	if err != nil {
		return &pb.UserResponse{Error: err.Error()}, nil
	}

	return &pb.UserResponse{
		User: s.modelUserToProto(user),
	}, nil
}

func (s *CardServiceServer) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var protoUsers []*pb.User
	for _, user := range users {
		protoUsers = append(protoUsers, s.modelUserToProto(user))
	}

	return &pb.GetAllUsersResponse{Users: protoUsers}, nil
}

func (s *CardServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	user := &models.User{
		ID:        int(req.Id),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	if req.Phone != "" {
		user.Phone = sql.NullString{String: req.Phone, Valid: true}
	}

	if req.BirthDate != "" {
		birthDate, err := time.Parse("2006-01-02", req.BirthDate)
		if err != nil {
			return &pb.UserResponse{Error: "Invalid birth date format"}, nil
		}
		user.BirthDate = sql.NullTime{Time: birthDate, Valid: true}
	}

	updatedUser, err := s.userRepo.Update(user)
	if err != nil {
		return &pb.UserResponse{Error: err.Error()}, nil
	}

	return &pb.UserResponse{
		User: s.modelUserToProto(updatedUser),
	}, nil
}

func (s *CardServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := s.userRepo.Delete(int(req.Id))
	if err != nil {
		return &pb.DeleteUserResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.DeleteUserResponse{
		Success: true,
		Message: "User deleted successfully",
	}, nil
}

// Операции с картами
func (s *CardServiceServer) CreateCard(ctx context.Context, req *pb.CreateCardRequest) (*pb.CardResponse, error) {
	issueDate, err := time.Parse("2006-01-02", req.IssueDate)
	if err != nil {
		return &pb.CardResponse{Error: "Invalid issue date format"}, nil
	}

	expiryDate, err := time.Parse("2006-01-02", req.ExpiryDate)
	if err != nil {
		return &pb.CardResponse{Error: "Invalid expiry date format"}, nil
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
		return &pb.CardResponse{Error: err.Error()}, nil
	}

	return &pb.CardResponse{
		Card: s.modelCardToProto(createdCard),
	}, nil
}

func (s *CardServiceServer) GetCard(ctx context.Context, req *pb.GetCardRequest) (*pb.CardResponse, error) {
	card, err := s.cardRepo.GetByID(int(req.Id))
	if err != nil {
		return &pb.CardResponse{Error: err.Error()}, nil
	}

	return &pb.CardResponse{
		Card: s.modelCardToProto(card),
	}, nil
}

func (s *CardServiceServer) GetAllCards(ctx context.Context, req *pb.GetAllCardsRequest) (*pb.GetAllCardsResponse, error) {
	cards, err := s.cardRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var protoCards []*pb.Card
	for _, card := range cards {
		protoCards = append(protoCards, s.modelCardToProto(card))
	}

	return &pb.GetAllCardsResponse{Cards: protoCards}, nil
}

func (s *CardServiceServer) DeleteCard(ctx context.Context, req *pb.DeleteCardRequest) (*pb.DeleteCardResponse, error) {
	err := s.cardRepo.Delete(int(req.Id))
	if err != nil {
		return &pb.DeleteCardResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.DeleteCardResponse{
		Success: true,
		Message: "Card deleted successfully",
	}, nil
}

func (s *CardServiceServer) GenerateCard(ctx context.Context, req *pb.GenerateCardRequest) (*pb.CardResponse, error) {
	card, err := s.cardRepo.GenerateCard(int(req.UserId), int(req.OperatorId), req.Balance)
	if err != nil {
		return &pb.CardResponse{Error: err.Error()}, nil
	}

	return &pb.CardResponse{
		Card: s.modelCardToProto(card),
	}, nil
}

// Вспомогательные методы для конвертации
func (s *CardServiceServer) modelUserToProto(user *models.User) *pb.User {
	protoUser := &pb.User{
		Id:        int32(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}

	if user.Phone.Valid {
		protoUser.Phone = user.Phone.String
	}

	if user.BirthDate.Valid {
		protoUser.BirthDate = user.BirthDate.Time.Format("2006-01-02")
	}

	return protoUser
}

func (s *CardServiceServer) modelCardToProto(card *models.Card) *pb.Card {
	return &pb.Card{
		Id:                 int32(card.ID),
		UserId:             int32(card.UserID),
		CardNumber:         card.CardNumber,
		OperatorName:       card.OperatorName,
		OperatorCode:       card.OperatorCode,
		IssueDate:          card.IssueDate.Format("2006-01-02"),
		ExpiryDate:         card.ExpiryDate.Format("2006-01-02"),
		IsActive:           card.IsActive,
		Balance:            card.Balance,
		CreatedAt:          timestamppb.New(card.CreatedAt),
		CardholderFullName: card.CardholderFullName,
	}
}
