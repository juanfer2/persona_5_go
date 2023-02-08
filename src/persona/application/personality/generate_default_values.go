package persona_application

import (
	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
	persona_repository "github.com/juanfer2/persona_5/src/persona/domain/repositories"
)

type GenerateDefaultValues struct {
	repository persona_repository.PersonalityRepository
}

func NewGenerateDefaultValues(
	repository persona_repository.PersonalityRepository,
) *GenerateDefaultValues {
	return &GenerateDefaultValues{
		repository: repository,
	}
}

func (generateDefaultValues GenerateDefaultValues) Call() {
	listArcana := []persona_entities.Personality{
		{Name: "Gloomy"},
		{Name: "Irritable"},
		{Name: "Timid"},
		{Name: "Upbeat"},
		{Name: "Vague"},
		{Name: "Serious"},
		{Name: "Kind"},
		{Name: "Funny"},
	}

	generateDefaultValues.repository.CreateInBatches(listArcana)
}
