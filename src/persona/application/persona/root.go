package persona_application

import (
	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
	persona_repository "github.com/juanfer2/persona_5/src/persona/domain/repositories"
)

type PersonaUseCase struct {
	repository persona_repository.PersonaRepository
}

func NewPersonaUseCase(repository persona_repository.PersonaRepository) *PersonaUseCase {
	personaUseCase := PersonaUseCase{repository: repository}

	return &personaUseCase
}

func (personaUseCase *PersonaUseCase) Create(persona persona_entities.Persona) persona_entities.Persona {
	return personaUseCase.repository.Create(persona)
}

func (personaUseCase *PersonaUseCase) All() []persona_entities.Persona {
	return personaUseCase.repository.All()
}

func (personaUseCase *PersonaUseCase) FindById(id int) persona_entities.Persona {
	return personaUseCase.repository.FindBy(id)
}

func (personaUseCase *PersonaUseCase) DestroyById(id int) persona_entities.Persona {
	persona := personaUseCase.repository.FindBy(id)
	personaUseCase.repository.Delete(persona)
	return persona
}

func (personaUseCase *PersonaUseCase) CreateMultiplesRecords(personaList []persona_entities.Persona) {
	personaUseCase.repository.CreateInBatches(personaList)
}
