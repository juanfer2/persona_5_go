package persona_application

import (
	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
	persona_repository "github.com/juanfer2/persona_5/src/persona/domain/repositories"
)

type ArcanaUseCase struct {
	repository persona_repository.ArcanaRepository
}

func NewArcanaUseCase(repository persona_repository.ArcanaRepository) *ArcanaUseCase {
	arcanaUseCase := ArcanaUseCase{repository: repository}

	return &arcanaUseCase
}

func (arcanaUseCase *ArcanaUseCase) Create(arcana persona_entities.Arcana) persona_entities.Arcana {
	return arcanaUseCase.repository.Create(arcana)
}

func (arcanaUseCase *ArcanaUseCase) All() []persona_entities.Arcana {
	return arcanaUseCase.repository.All()
}

func (arcanaUseCase *ArcanaUseCase) FindById(id int) persona_entities.Arcana {
	return arcanaUseCase.repository.FindBy(id)
}

func (arcanaUseCase *ArcanaUseCase) DestroyById(id int) persona_entities.Arcana {
	arcana := arcanaUseCase.repository.FindBy(id)
	arcanaUseCase.repository.Delete(arcana)
	return arcana
}

func (arcanaUseCase *ArcanaUseCase) CreateMultiplesRecords(arcanaList []persona_entities.Arcana) {
	arcanaUseCase.repository.CreateInBatches(arcanaList)
}
