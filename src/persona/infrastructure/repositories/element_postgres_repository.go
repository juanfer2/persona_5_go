package persona_repositories

import (
	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
	persona_repository "github.com/juanfer2/persona_5/src/persona/domain/repositories"
	"github.com/juanfer2/persona_5/src/shared/infrastructure/persistence/postgres"
)

type ElementPosgrestRepository struct {
	postgres.PostgresRepository[persona_entities.Element]
}

func NewElementPostgresRepository() persona_repository.ElementRepository {
	repo := ElementPosgrestRepository{}
	repo.PostgresClient = postgres.CreateClientFactory()

	return &repo
}
