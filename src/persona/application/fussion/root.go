package persona_application_fussion

import (
	"math"

	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
	persona_repository "github.com/juanfer2/persona_5/src/persona/domain/repositories"
)

type CalculatorUseCase struct {
	personaRepository persona_repository.PersonaRepository
}

func NewCalculatorUseCase(
	personaRepository persona_repository.PersonaRepository,
) *CalculatorUseCase {
	return &CalculatorUseCase{personaRepository: personaRepository}
}

func (c *CalculatorUseCase) Fusion(
	personaIDFirst int,
	personaIDSecond int,
) persona_entities.Persona {
	personaFirst := c.personaRepository.FindBy(personaIDFirst)
	personaSecond := c.personaRepository.FindBy(personaIDSecond)

	if personaFirst.Arcana.ID == personaSecond.Arcana.ID {
		return c.fussionArcana(personaFirst, personaSecond)
	} else {
		return c.fussionNormal(personaFirst, personaSecond)
	}
}

func (c *CalculatorUseCase) fussionArcana(
	personaFirst persona_entities.Persona,
	personaSecond persona_entities.Persona,
) persona_entities.Persona {
	level := 1 + math.Floor(float64(personaFirst.Level+personaSecond.Level)/2)
	persona := c.personaRepository.FindFussionArcana(level, personaFirst.ID, personaSecond.ID)

	return persona
}

func (c *CalculatorUseCase) fussionNormal(
	personaFirst persona_entities.Persona,
	personaSecond persona_entities.Persona,
) persona_entities.Persona {
	level := 1 + math.Floor(float64(personaFirst.Level+personaSecond.Level)/2)
	persona := c.personaRepository.FindFussionNormal(level, personaFirst.ID, personaSecond.ID)

	return persona
}

// func (c *CalculatorUseCase) fussionSpecial(
// 	personaFirst persona_entities.Persona,
// 	personaSecond persona_entities.Persona,
// ) {
//
// }
