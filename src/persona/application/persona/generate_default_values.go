package persona_application

import (
	"fmt"

	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
	persona_repository "github.com/juanfer2/persona_5/src/persona/domain/repositories"
)

type GenerateDefaultValues struct {
	personaRepository     persona_repository.PersonaRepository
	arcanaRepository      persona_repository.ArcanaRepository
	personalityRepository persona_repository.PersonalityRepository
	statRepository        persona_repository.StatRepository
}

func NewGenerateDefaultValues(
	personaRepository persona_repository.PersonaRepository,
	arcanaRepository persona_repository.ArcanaRepository,
	personalityRepository persona_repository.PersonalityRepository,
	statRepository persona_repository.StatRepository,
) *GenerateDefaultValues {
	return &GenerateDefaultValues{
		personaRepository:     personaRepository,
		arcanaRepository:      arcanaRepository,
		personalityRepository: personalityRepository,
		statRepository:        statRepository,
	}
}

func (generateDefaultValues GenerateDefaultValues) Call() {
	timid := generateDefaultValues.findPersonalityByName("Timid").ID
	listArcana := []persona_entities.Persona{
		{Name: "Ame-no-Uzume", ImageUrl: "ame-no-uzume.png", Level: 29,
			ArcanaID: generateDefaultValues.findArcanaByName("Lovers").ID,
			PersonaStats: []persona_entities.PersonaStat{
				{Value: 15, StatID: generateDefaultValues.findStatByName("Strength").ID},
				{Value: 15, StatID: generateDefaultValues.findStatByName("Magic").ID},
				{Value: 15, StatID: generateDefaultValues.findStatByName("Endurance").ID},
				{Value: 15, StatID: generateDefaultValues.findStatByName("Agility").ID},
				{Value: 15, StatID: generateDefaultValues.findStatByName("Luck").ID},
			},
		},
		{Name: "Ananta", ImageUrl: "ananta.png", Level: 43,
			ArcanaID: generateDefaultValues.findArcanaByName("Star").ID},
		{Name: "Andras", ImageUrl: "andras.png", Level: 10,
			PersonalityID: &timid,
			ArcanaID:      generateDefaultValues.findArcanaByName("Devil").ID},
		{Name: "Angel", ImageUrl: "angel.png", Level: 12,
			ArcanaID: generateDefaultValues.findArcanaByName("Justice").ID},
	}

	generateDefaultValues.personaRepository.CreateInBatches(listArcana)

	p := generateDefaultValues.personaRepository.All()

	for _, v := range p {
		fmt.Println(v.Name)
		fmt.Println(v.Arcana.Name)
		fmt.Println(v.Personality.Name)
	}
}

func (generateDefaultValues *GenerateDefaultValues) findArcanaByName(
	name string,
) persona_entities.Arcana {
	arcana := generateDefaultValues.arcanaRepository.WhereBy("name=?", name)
	return arcana
}

func (generateDefaultValues *GenerateDefaultValues) findPersonalityByName(
	name string,
) persona_entities.Personality {
	personality := generateDefaultValues.personalityRepository.WhereBy("name=?", name)
	return personality
}

func (generateDefaultValues *GenerateDefaultValues) findStatByName(
	name string,
) persona_entities.Stat {
	stat := generateDefaultValues.statRepository.WhereBy("name=?", name)
	return stat
}
