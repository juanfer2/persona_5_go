package persona_entities

import (
	"time"

	"gorm.io/gorm"
)

type Persona struct {
	gorm.Model
	ID       int
	Name     string
	Level    int
	Rare     bool
	ImageUrl string

	ArcanaID int
	Arcana   Arcana

	PersonalityID *int
	Personality   Personality

	Elems        []Elem
	PersonaStats []PersonaStat

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Persona) IsRare() bool {
	return p.Rare
}
