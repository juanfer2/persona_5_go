package persona_repositories

import (
	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
	persona_repository "github.com/juanfer2/persona_5/src/persona/domain/repositories"
	"github.com/juanfer2/persona_5/src/shared/infrastructure/persistence/postgres"
)

type StatPostgresRepository struct {
	postgres.PostgresRepository[persona_entities.Stat]
}

func NewStatPostgresRepository() persona_repository.StatRepository {
	repo := StatPostgresRepository{}
	repo.PostgresClient = postgres.CreateClientFactory()

	return &repo
}
