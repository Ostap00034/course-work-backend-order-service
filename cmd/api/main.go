package main

import (
	"log"
	"net"
	"os"

	orderpbv1 "github.com/Ostap00034/course-work-backend-api-specs/gen/go/order/v1"
	userpbv1 "github.com/Ostap00034/course-work-backend-api-specs/gen/go/user/v1"

	"github.com/Ostap00034/course-work-backend-order-service/db"
	order "github.com/Ostap00034/course-work-backend-order-service/internal"
	"github.com/joho/godotenv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	dbString, ok := os.LookupEnv("DB_CONN_STRING")
	if !ok {
		log.Fatal("DB_CONN_STRING is not set")
	}
	client := db.NewClient(dbString)
	defer client.Close()

	repo := order.NewRepo(client)
	svc := order.NewService(repo)

	userAddr, ok := os.LookupEnv("USER_SERVICE_ADDR")
	if !ok {
		log.Fatal("USER_SERVICE_ADDR is not set")
	}
	userConn, err := grpc.NewClient(
		userAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial UserService: %v", err)
	}
	defer userConn.Close()
	userSvc := userpbv1.NewUserServiceClient(userConn)

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcSrv := grpc.NewServer()
	srv := order.NewServer(svc, userSvc)
	orderpbv1.RegisterOrderServiceServer(grpcSrv, srv)

	log.Println("OrderService is listening on :50054")
	log.Fatal(grpcSrv.Serve(lis))
}
