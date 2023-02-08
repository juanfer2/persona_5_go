package persona_repository

import (
	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
)

type DefaultRepository[T any] interface {
	FindBy(id int) T
	Create(newRecord T) T
	All() []T
	Delete(record T) T
	WhereBy(query any, arg ...any) T
	// FindWithAssociation(associations []string, query any, arg ...any) T
	// WhereWithAssociation(associations []string, query any, arg ...any) []T
	// Join(relations ...string) *gorm.DB
	CreateInBatches(records []T)
}

type CharacterRepository interface {
	DefaultRepository[persona_entities.Character]
}

type ArcanaRepository interface {
	DefaultRepository[persona_entities.Arcana]
}

type ElementRepository interface {
	DefaultRepository[persona_entities.Element]
}

type ElemRepository interface {
	DefaultRepository[persona_entities.Elem]
}

type PersonalityRepository interface {
	DefaultRepository[persona_entities.Personality]
}

type PersonaRepository interface {
	DefaultRepository[persona_entities.Persona]
	FindFussionNormal(level float64, personaIDFirst int, personaIDSecond int) persona_entities.Persona
	FindFussionArcana(level float64, personaIDFirst int, personaIDSecond int) persona_entities.Persona
}

type StatRepository interface {
	DefaultRepository[persona_entities.Stat]
}

type PersonaStatRepository interface {
	DefaultRepository[persona_entities.PersonaStat]
}
