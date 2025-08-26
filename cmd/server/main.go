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
	// Connect to the database
	db, err := database.NewDB(DbHost, DbPort, DbUser, DbPassword, DbName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create repositories
	userRepo := repository.NewUserRepository(db)
	cardRepo := repository.NewCardRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	grpcServer := grpc.NewServer()
	pb.RegisterCardServiceServer(grpcServer, service.NewCardServiceServer(userRepo, cardRepo, transactionRepo))

	// Start listening on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server starting on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
