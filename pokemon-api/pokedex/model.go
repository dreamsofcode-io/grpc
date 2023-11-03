package pokedex

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/dreamsofcode-io/grpc/pokemon-api/pb"
)

type Pokemon struct {
	ID        int32
	Name      string
	Types     []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func pokemonToResponse(p Pokemon) pb.Pokemon {
	types := make([]pb.Type, len(p.Types))
	for i, t := range p.Types {
		types[i] = pb.Type(pb.Type_value[t])
	}

	return pb.Pokemon{
		Id:        p.ID,
		Name:      p.Name,
		Type:      types,
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt),
	}
}
