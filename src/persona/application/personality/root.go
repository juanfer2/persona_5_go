package persona_application

import (
	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
	persona_repository "github.com/juanfer2/persona_5/src/persona/domain/repositories"
)

type PersonalityUseCase struct {
	repository persona_repository.PersonalityRepository
}

func (personalityUseCase *PersonalityUseCase) Create(
	name string, imageUrl string,
) persona_entities.Personality {
	return personalityUseCase.repository.Create(persona_entities.Personality{Name: name})
}

func (personalityUseCase *PersonalityUseCase) All() []persona_entities.Personality {
	return personalityUseCase.repository.All()
}

func (personalityUseCase *PersonalityUseCase) FindById(id int) persona_entities.Personality {
	return personalityUseCase.repository.FindBy(id)
}

func (personalityUseCase *PersonalityUseCase) DestroyById(id int) persona_entities.Personality {
	Personality := personalityUseCase.repository.FindBy(id)
	personalityUseCase.repository.Delete(Personality)
	return Personality
}

func (personalityUseCase *PersonalityUseCase) CreateMultiplesRecords(
	PersonalityList []persona_entities.Personality,
) {
	personalityUseCase.repository.CreateInBatches(PersonalityList)
}
