package persona_application

import (
	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
	persona_repository "github.com/juanfer2/persona_5/src/persona/domain/repositories"
)

type ElementUseCase struct {
	repository persona_repository.ElementRepository
}

func (elementUseCase *ElementUseCase) Create(
	name string, imageUrl string,
) persona_entities.Element {
	return elementUseCase.repository.Create(persona_entities.Element{Name: name, ImageUrl: imageUrl})
}

func (elementUseCase *ElementUseCase) All() []persona_entities.Element {
	return elementUseCase.repository.All()
}

func (elementUseCase *ElementUseCase) FindById(id int) persona_entities.Element {
	return elementUseCase.repository.FindBy(id)
}

func (elementUseCase *ElementUseCase) DestroyById(id int) persona_entities.Element {
	element := elementUseCase.repository.FindBy(id)
	elementUseCase.repository.Delete(element)
	return element
}

func (elementUseCase *ElementUseCase) CreateMultiplesRecords(
	elementList []persona_entities.Element,
) {
	elementUseCase.repository.CreateInBatches(elementList)
}
