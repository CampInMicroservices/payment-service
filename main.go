package main

import (
	"log"
	"net"

	"payment-service/api"
	"payment-service/config"
	"payment-service/db"
	"payment-service/gapi"
	"payment-service/proto"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"google.golang.org/grpc"
)

func main() {

	// Load configuration settings
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Connect to the database
	store, err := db.Connect(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Run DB migration
	runDBMigration(config.MigrationURL, config.DBSource)

	// Create a server and setup routes
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Failed to create a server: ", err)
	}

	// GRPC server (concurrently)
	go func() {
		lis, err := net.Listen("tcp", config.GRPCAddress)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		s := gapi.NewGrpcServer(config, store)
		grpcServer := grpc.NewServer()

		proto.RegisterPaymentServiceServer(grpcServer, &s)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	// Start a server
	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("Failed to start a server: ", err)
	}
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("Cannot create new migrate instance", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Failed to run migrate up", err)
	}

	log.Println("Db migrated successfully")
}
