package servers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	persona_routes "github.com/juanfer2/persona_5/src/persona/infrastructure/routes"
	"github.com/juanfer2/persona_5/src/shared/infrastructure/routes"
)

func StartServerApp() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:   "${pid} ${status} - ${method} ${path}\n",
		TimeZone: "America/Bogotá",
		Done: func(c *fiber.Ctx, logString []byte) {
			fmt.Sprintf("%s - %d", c.Request().RequestURI(), c.Response().StatusCode())
		},
	}))

	app.Use(cors.New())
	app.Static("/images", "./tmp/images")

	routes.SetupRoutes(app)
	persona_routes.SetUpRoutes(app)

	app.Listen(":4000")
}
