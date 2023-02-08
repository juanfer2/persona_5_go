package package_seed

import (
	"github.com/fatih/color"

	arcana_use_case "github.com/juanfer2/persona_5/src/persona/application/arcana"
	element_use_case "github.com/juanfer2/persona_5/src/persona/application/elements"
	persona_use_case "github.com/juanfer2/persona_5/src/persona/application/persona"
	personality_use_case "github.com/juanfer2/persona_5/src/persona/application/personality"
	stat_use_case "github.com/juanfer2/persona_5/src/persona/application/stat"
	persona_repositories "github.com/juanfer2/persona_5/src/persona/infrastructure/repositories"
)

func ExecCommand() {

	color.Yellow(">>>Run Arcanas 🚀 <<<")
	arcanaRepository := persona_repositories.NewArcanaPostgresRepository()
	arcanaService := arcana_use_case.NewGenerateDefaultValues(arcanaRepository)
	arcanaService.Call()

	color.Yellow(">>>Run Elements 🚀 <<<")
	elementRepository := persona_repositories.NewElementPostgresRepository()
	elementService := element_use_case.NewGenerateDefaultValues(elementRepository)
	elementService.Call()

	color.Yellow(">>>Run Personalities 🚀 <<<")
	personalityRepository := persona_repositories.NewPersonalityPostgresRepository()
	personalityService := personality_use_case.NewGenerateDefaultValues(personalityRepository)
	personalityService.Call()

	color.Yellow(">>>Run Stats 🚀 <<<")
	statRepository := persona_repositories.NewStatPostgresRepository()
	statService := stat_use_case.NewGenerateDefaultValues(statRepository)
	statService.Call()

	color.Yellow(">>>Run Personas 🚀 <<<")
	personaRepository := persona_repositories.NewPersonaPostgresRepository()
	personaService := persona_use_case.NewGenerateDefaultValues(personaRepository, arcanaRepository,
		personalityRepository, statRepository)
	personaService.Call()
}
