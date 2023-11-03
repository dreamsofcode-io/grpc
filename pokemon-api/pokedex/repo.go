package pokedex

import (
	"context"
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("pokemon not found")
var ErrAlreadyExists = errors.New("pokemon already exists")

type Repository struct {
	db  *sql.DB
	mem map[int32]Pokemon
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db:  db,
		mem: make(map[int32]Pokemon),
	}
}

func (p *Repository) Insert(ctx context.Context, pokemon Pokemon) error {
	if _, exists := p.mem[pokemon.ID]; exists {
		return ErrAlreadyExists
	}

	p.mem[pokemon.ID] = pokemon
	return nil
}

func (p *Repository) FindAll(ctx context.Context) ([]Pokemon, error) {
	res := make([]Pokemon, 0, len(p.mem))
	for _, pokemon := range p.mem {
		res = append(res, pokemon)
	}

	return res, nil
}

func (p *Repository) FindByID(ctx context.Context, id int32) (Pokemon, error) {
	pokemon, exists := p.mem[id]
	if !exists {
		return Pokemon{}, ErrNotFound
	}

	return pokemon, nil
}

func (p *Repository) Update(ctx context.Context, pokemon Pokemon) error {
	if _, exists := p.mem[pokemon.ID]; !exists {
		return ErrNotFound
	}

	p.mem[pokemon.ID] = pokemon

	return nil
}
