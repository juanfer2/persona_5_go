package persona_entities

import (
	"time"

	"gorm.io/gorm"
)

type Element struct {
	gorm.Model
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
