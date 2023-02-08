package persona_entities

import (
	"time"

	"gorm.io/gorm"
)

type Arcana struct {
	gorm.Model
	ID                int    `json:"id"`
	Name              string `json:"name"`
	ImageUrl          string `json:"image_url"`
	CharacterName     string `json:"character_name"`
	ImageCharacterUrl string `json:"image_character_url"`

	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt time.Time
}
