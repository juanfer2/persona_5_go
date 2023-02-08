package persona_routes

import (
	"github.com/gofiber/fiber/v2"

	persona_application "github.com/juanfer2/persona_5/src/persona/application/persona"
	persona_controller "github.com/juanfer2/persona_5/src/persona/infrastructure/controller"
	persona_repositories "github.com/juanfer2/persona_5/src/persona/infrastructure/repositories"
)

func SetUpRoutes(app *fiber.App) {
	personaRepository := persona_repositories.NewPersonaPostgresRepository()
	personaUseCase := persona_application.NewPersonaUseCase(personaRepository)
	personaController := persona_controller.NewPersonaController(*personaUseCase)

	app.Get("/personas", personaController.All)
}
