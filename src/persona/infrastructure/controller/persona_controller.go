package persona_controller

import (
	"github.com/gofiber/fiber/v2"
	persona_application "github.com/juanfer2/persona_5/src/persona/application/persona"
	persona_serializer "github.com/juanfer2/persona_5/src/persona/infrastructure/serializer"
)

type PersonaController struct {
	personaUseCase persona_application.PersonaUseCase
}

func NewPersonaController(personaUseCase persona_application.PersonaUseCase) *PersonaController {
	return &PersonaController{personaUseCase: personaUseCase}
}

func (personaController *PersonaController) All(c *fiber.Ctx) error {
	personas := personaController.personaUseCase.All()

	return c.JSON(persona_serializer.PersonasToSerializer(personas))
}
