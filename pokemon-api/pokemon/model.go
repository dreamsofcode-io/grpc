package pokemon

import (
	"time"
)

type Pokemon struct {
	ID        int32
	Name      string
	Types     []string
	CreatedAt time.Time
	UpdatedAt time.Time
}
