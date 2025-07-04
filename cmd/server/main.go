package main

import (
	"log"
	"net"

	pb "card_service/generated/proto"
	"card_service/internal/database"
	"card_service/internal/repository"
	"card_service/internal/service"

	"google.golang.org/grpc"
)

func main() {
	// Подключение к базе данных
	db, err := database.NewDB(DbHost, DbPort, DbUser, DbPassword, DbName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Создание репозиториев
	userRepo := repository.NewUserRepository(db)
	cardRepo := repository.NewCardRepository(db)

	// Создание gRPC сервера
	grpcServer := grpc.NewServer()
	cardService := service.NewCardServiceServer(userRepo, cardRepo)
	pb.RegisterCardServiceServer(grpcServer, cardService)

	// Запуск сервера
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server starting on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
