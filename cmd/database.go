package cmd

import (
	"log"

	"github.com/fatih/color"
	package_seed "github.com/juanfer2/persona_5/src/persona/infrastructure/persitence/seed"
	"github.com/juanfer2/persona_5/src/shared/infrastructure/persistence/postgres"
	"github.com/spf13/cobra"
)

var createDatBaseCmd = &cobra.Command{
	Use:   "db:create",
	Short: "Create database",
	Long:  `Create database`,
	Run:   createDB,
}

func createDB(cmd *cobra.Command, args []string) {
	configClient := postgres.CreateConfigClient()
	postgresClient := postgres.PostgresClient{ConfigClient: configClient}
	postgresClient.CreateDataBase()
}

var setUpDbCmd = &cobra.Command{
	Use:   "db:setup",
	Short: "Setup database",
	Long:  `Setup database`,
	Run:   setUpDb,
}

func setUpDb(cmd *cobra.Command, args []string) {
	color.Blue(">>>Create DB 🚀 <<<")
	configClient := postgres.CreateConfigClient()
	postgresClient := postgres.PostgresClient{ConfigClient: configClient}
	postgresClient.CreateDataBase()

	color.Blue(">>>RunMigration 🚀 <<<")
	m := postgres.NewMigrationClient()
	err := m.Migrate.Up()

	if err != nil {
		color.Red(err.Error())
		log.Fatal(err)
	}

	color.Blue(">>>RunSeed 🚀 <<<")
	package_seed.ExecCommand()
}
