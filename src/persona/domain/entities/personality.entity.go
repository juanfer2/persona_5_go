package persona_entities

import (
	"time"

	"gorm.io/gorm"
)

type Personality struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`

	Personas []Persona

	CreatedAt time.Time
	UpdatedAt time.Time
}

/*
type Personality struct {
	ID              int
	PersonalityType PersonalityType
	Likes           []Like
	Dislikes        []Dislike
}

type PersonalityType struct {
	ID   int
	Name string
}

type Like struct {
	ID                int
	PersonalityTypeID int
	PersonalityType   PersonalityType
}

type Dislike struct {
	ID                int
	PersonalityTypeID int
	PersonalityType   PersonalityType
}
*/
/*

Type	Likes	Dislikes
Gloomy	Vague	Serious, Funny
Irritable	Serious	Vague, Kind
Timid	Kind	Vague, Funny
Upbeat	Funny	Serious, Vague
*/
