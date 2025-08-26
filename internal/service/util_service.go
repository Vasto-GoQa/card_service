package service

import (
	pb "card_service/generated/proto"
	"card_service/internal/models"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Method to convert user model to protobuf messages
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

// Method to convert card model to protobuf messages
func (s *CardServiceServer) modelCardToProto(card *models.Card, operator *models.CardOperator) *pb.Card {
	return &pb.Card{
		Id:         int32(card.ID),
		UserId:     int32(card.UserID),
		CardNumber: card.CardNumber,
		Operator:   s.modelOperatorToProto(operator),
		IssueDate:  card.IssueDate.Format("2006-01-02"),
		ExpiryDate: card.ExpiryDate.Format("2006-01-02"),
		IsActive:   card.IsActive,
		Balance:    card.Balance,
		CreatedAt:  timestamppb.New(card.CreatedAt),
	}
}

// Method to convert operator model to protobuf messages
func (s *CardServiceServer) modelOperatorToProto(operator *models.CardOperator) *pb.CardOperator {
	if operator == nil {
		return nil
	}
	return &pb.CardOperator{
		Id:   int32(operator.ID),
		Name: operator.Name,
		Code: operator.Code,
	}
}

// Method to convert transaction model to protobuf messages
func (s *CardServiceServer) modelTransactionToProto(transaction *models.Transaction) *pb.Transaction {
	return &pb.Transaction{
		Id:         int32(transaction.ID),
		FromCardId: int32(transaction.FromCardID),
		ToCardId:   int32(transaction.ToCardID),
		Amount:     transaction.Amount,
		CreatedAt:  timestamppb.New(transaction.CreatedAt),
	}
}

// Method to convert array transaction model to protobuf messages
func (s *CardServiceServer) modelTransactionsToProto(transactions []*models.Transaction) []*pb.Transaction {
	protoTransactions := make([]*pb.Transaction, len(transactions))
	for i, t := range transactions {
		protoTransactions[i] = s.modelTransactionToProto(t)
	}
	return protoTransactions
}
