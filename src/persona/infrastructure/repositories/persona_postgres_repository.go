package persona_repositories

import (
	"fmt"

	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
	persona_repository "github.com/juanfer2/persona_5/src/persona/domain/repositories"
	"github.com/juanfer2/persona_5/src/shared/infrastructure/persistence/postgres"
)

type PersonaPostgresRepository struct {
	postgres.PostgresRepository[persona_entities.Persona]
}

func NewPersonaPostgresRepository() persona_repository.PersonaRepository {
	repo := PersonaPostgresRepository{}
	repo.PostgresClient = postgres.CreateClientFactory()

	return &repo
}

func (personaPostgresRepository *PersonaPostgresRepository) All() []persona_entities.Persona {
	var personas []persona_entities.Persona
	err := personaPostgresRepository.PostgresRepository.Preload("PersonaStats", "PersonaStats.Stat", "Personality", "Arcana").
		//Join("Arcana", "Personality").
		// Association("PersonaStats").
		Find(&personas).Error
	if err != nil {
		fmt.Println(err)
	}

	return personas
}

func (personaPostgresRepository *PersonaPostgresRepository) FindFussionArcana(
	level float64, personaIDFirst int, personaIDSecond int,
) (persona persona_entities.Persona) {
	personaPostgresRepository.PostgresRepository.
		Where(
			"level <= ? AND special = false AND rare = false AND id NOT IN (?, ?) ? and",
			level, personaIDFirst, personaIDSecond,
		).First(persona)

	return
}

func (personaPostgresRepository *PersonaPostgresRepository) FindFussionNormal(
	level float64, personaIDFirst int, personaIDSecond int,
) (persona persona_entities.Persona) {
	personaPostgresRepository.PostgresRepository.
		Where(
			"level >= ? AND special = false AND rare = false AND id NOT IN (?, ?) ? and",
			level, personaIDFirst, personaIDSecond,
		).First(persona)

	return
}
