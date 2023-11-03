package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

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

	repository := pokedex.NewRepository(nil)
	server := pokedex.NewServer(repository)

	pb.RegisterPokedexServer(s, server)
	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
