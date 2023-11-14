package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/dreamsofcode-io/grpc/pokemon-api/database"
	"github.com/dreamsofcode-io/grpc/pokemon-api/pb"
	"github.com/dreamsofcode-io/grpc/pokemon-api/pokedex"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("failed to create listener:", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pool, err := database.Connect(context.Background())
	if err != nil {
		log.Fatalln("failed to connect to database:", err)
	}

	repository := pokedex.NewRepository(pool)
	server := pokedex.NewServer(repository)

	pb.RegisterPokedexServer(s, server)
	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
