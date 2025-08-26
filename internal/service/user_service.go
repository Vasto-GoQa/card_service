package service

import (
	pb "card_service/generated/proto"
	"card_service/internal/models"
	"context"
	"database/sql"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CardServiceServer) CreateUser(_ context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
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
			return nil, status.Errorf(codes.InvalidArgument, "Invalid birth date format")
		}
		user.BirthDate = sql.NullTime{Time: birthDate, Valid: true}
	}

	createdUser, err := s.userRepo.Create(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create user: %v", err)
	}

	return s.modelUserToProto(createdUser), nil
}

func (s *CardServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user, err := s.userRepo.GetByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get user: %v", err)
	}

	return s.modelUserToProto(user), nil
}

func (s *CardServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
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
			return nil, status.Errorf(codes.InvalidArgument, "Invalid birth date format")
		}
		user.BirthDate = sql.NullTime{Time: birthDate, Valid: true}
	}

	updatedUser, err := s.userRepo.Update(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user: %v", err)
	}

	return s.modelUserToProto(updatedUser), nil
}

func (s *CardServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := s.userRepo.Delete(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete user: %v", err)
	}

	return &pb.DeleteUserResponse{
		Success: true,
		Message: "User deleted successfully",
	}, nil
}
