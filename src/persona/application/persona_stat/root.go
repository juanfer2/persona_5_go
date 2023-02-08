package persona_application

import (
	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
	persona_repository "github.com/juanfer2/persona_5/src/persona/domain/repositories"
)

type ArcanaUseCase struct {
	repository persona_repository.StatRepository
}

func NewArcanaUseCase(repository persona_repository.StatRepository) *ArcanaUseCase {
	arcanaUseCase := ArcanaUseCase{repository: repository}

	return &arcanaUseCase
}

func (arcanaUseCase *ArcanaUseCase) Create(stat persona_entities.Stat) persona_entities.Stat {
	return arcanaUseCase.repository.Create(stat)
}

func (arcanaUseCase *ArcanaUseCase) All() []persona_entities.Stat {
	return arcanaUseCase.repository.All()
}

func (arcanaUseCase *ArcanaUseCase) FindById(id int) persona_entities.Stat {
	return arcanaUseCase.repository.FindBy(id)
}

func (arcanaUseCase *ArcanaUseCase) DestroyById(id int) persona_entities.Stat {
	stat := arcanaUseCase.repository.FindBy(id)
	arcanaUseCase.repository.Delete(stat)
	return stat
}

func (arcanaUseCase *ArcanaUseCase) CreateMultiplesRecords(stats []persona_entities.Stat) {
	arcanaUseCase.repository.CreateInBatches(stats)
}
