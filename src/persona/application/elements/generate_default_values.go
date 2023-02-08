package persona_application

import (
	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
	persona_repository "github.com/juanfer2/persona_5/src/persona/domain/repositories"
)

type GenerateDefaultValues struct {
	repository persona_repository.ElementRepository
}

func NewGenerateDefaultValues(
	repository persona_repository.ElementRepository,
) *GenerateDefaultValues {
	return &GenerateDefaultValues{
		repository: repository,
	}
}

func (generateDefaultValues GenerateDefaultValues) Call() {
	listArcana := []persona_entities.Element{
		{Name: "Phys", ImageUrl: "phys"},
		{Name: "Gun", ImageUrl: "gun"},
		{Name: "Fire", ImageUrl: "fire"},
		{Name: "Ice", ImageUrl: "ice"},
		{Name: "Electric", ImageUrl: "electric"},
		{Name: "Wind", ImageUrl: "wind"},
		{Name: "Psychokinesis", ImageUrl: "psychokinesis"},
		{Name: "Nuclear", ImageUrl: "nuclear"},
		{Name: "Bless", ImageUrl: "bless"},
		{Name: "Curse", ImageUrl: "curse"},
		{Name: "Almighty", ImageUrl: "almighty"},
		{Name: "Ailment Skills", ImageUrl: "ailment_skills"},
		{Name: "Healing Skills", ImageUrl: "healing_skills"},
		{Name: "Support Skills", ImageUrl: "support_skills"},
	}

	generateDefaultValues.repository.CreateInBatches(listArcana)
}
