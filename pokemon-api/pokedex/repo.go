package pokedex

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("pokemon not found")
var ErrAlreadyExists = errors.New("pokemon already exists")

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (p *Repository) Insert(ctx context.Context, pokemon Pokemon) error {
	_, err := p.db.Exec(
		ctx,
		"INSERT INTO pokemon (id, name, types, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		pokemon.ID,
		pokemon.Name,
		pokemon.Types,
		pokemon.CreatedAt,
		pokemon.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to insert pokemon: %w", err)
	}

	return nil
}

func (p *Repository) FindAll(ctx context.Context) ([]Pokemon, error) {
	rows, err := p.db.Query(ctx, "SELECT id, name, types, created_at, updated_at FROM pokemon")
	if err != nil {
		return nil, fmt.Errorf("failed to find all pokemon: %w", err)
	}

	defer rows.Close()
	res := make([]Pokemon, 0)

	for rows.Next() {
		var pokemon Pokemon

		err := rows.Scan(
			&pokemon.ID,
			&pokemon.Name,
			&pokemon.Types,
			&pokemon.CreatedAt,
			&pokemon.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan pokemon: %w", err)
		}

		res = append(res, pokemon)
	}

	return res, nil
}

func (p *Repository) FindByID(ctx context.Context, id int32) (Pokemon, error) {
	res := Pokemon{}

	err := p.db.QueryRow(
		ctx,
		"SELECT id, name, types, created_at, updated_at FROM pokemon WHERE id = $1",
		id,
	).Scan(
		&res.ID,
		&res.Name,
		&res.Types,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return Pokemon{}, ErrNotFound
	} else if err != nil {
		return Pokemon{}, fmt.Errorf("failed to find pokemon by id: %w", err)
	}

	return res, nil
}

func (p *Repository) Update(ctx context.Context, pokemon Pokemon) error {
	res, err := p.db.Exec(
		ctx,
		"UPDATE pokemon SET name = $1, types = $2, updated_at = $3 WHERE id = $4",
		pokemon.Name,
		pokemon.Types,
		pokemon.UpdatedAt,
		pokemon.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update pokemon: %w", err)
	}

	if res.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
