package persona_entities

import (
	"time"

	"gorm.io/gorm"
)

type PersonaStat struct {
	gorm.Model
	ID        int     `json:"id"`
	PersonaID int     `json:"personaId"`
	Persona   Persona `json:"persona"`
	StatID    int     `json:"statId"`
	Stat      Stat    `json:"stat"`
	Value     int     `json:"value"`

	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt time.Time
}
