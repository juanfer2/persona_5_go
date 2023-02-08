package persona_application

import (
	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
	persona_repository "github.com/juanfer2/persona_5/src/persona/domain/repositories"
)

type GenerateDefaultValues struct {
	repository persona_repository.StatRepository
}

func NewGenerateDefaultValues(
	repository persona_repository.StatRepository,
) *GenerateDefaultValues {
	return &GenerateDefaultValues{
		repository: repository,
	}
}

func (generateDefaultValues GenerateDefaultValues) Call() {
	stats := []persona_entities.Stat{
		{Name: "Strength"},
		{Name: "Magic"},
		{Name: "Endurance"},
		{Name: "Agility"},
		{Name: "Luck"},
	}

	generateDefaultValues.repository.CreateInBatches(stats)
}
