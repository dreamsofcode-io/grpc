package pokedex

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/dreamsofcode-io/grpc/pokemon-api/pb"
)

type Server struct {
	repository *Repository
	pb.UnimplementedPokedexServer
}

func NewServer(repository *Repository) *Server {
	return &Server{
		repository: repository,
	}
}

func (s *Server) Create(
	ctx context.Context,
	req *pb.PokemonRequest,
) (*pb.Pokemon, error) {
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be 0")
	}

	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name cannot be empty")
	}

	if len(req.Type) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "type cannot be empty")
	}

	types := make([]string, len(req.Type))
	for i, t := range req.Type {
		types[i] = t.String()
	}

	now := time.Now()

	pokemon := Pokemon{
		ID:        req.Id,
		Name:      req.Name,
		Types:     types,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repository.Insert(ctx, pokemon); err != nil {
		return nil, fmt.Errorf("failed to insert pokemon: %w", err)
	}

	res := pokemonToResponse(pokemon)

	return &res, nil
}

func (s *Server) Read(
	ctx context.Context,
	req *pb.PokemonFilter,
) (*pb.PokemonListResponse, error) {
	pokemon, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find all pokemon: %w", err)
	}

	res := make([]*pb.Pokemon, len(pokemon))

	for i, pokemon := range pokemon {
		p := pokemonToResponse(pokemon)
		res[i] = &p
	}

	return &pb.PokemonListResponse{
		Pokemon: res,
	}, nil
}

func (s *Server) ReadOne(
	ctx context.Context,
	req *pb.PokemonID,
) (*pb.Pokemon, error) {
	pokemon, err := s.repository.FindByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to find pokemon")
	}

	res := pokemonToResponse(pokemon)

	return &res, nil
}

func (s *Server) Update(
	ctx context.Context,
	req *pb.PokemonUpdateRequest,
) (*pb.Pokemon, error) {
	p, err := s.repository.FindByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to find pokemon")
	}

	strTypes := make([]string, len(req.Type))
	for i, t := range req.Type {
		strTypes[i] = t.String()
	}

	p.UpdatedAt = time.Now()
	p.Name = req.Name
	p.Types = strTypes

	err = s.repository.Update(ctx, p)
	if errors.Is(err, ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, "failed to update pokemon")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update pokemon")
	}

	res := pokemonToResponse(p)

	return &res, nil
}

func (s *Server) Delete(
	ctx context.Context,
	req *pb.PokemonID,
) (*emptypb.Empty, error) {
	return nil, nil
}
