package persona_cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	package_seed "github.com/juanfer2/persona_5/src/persona/infrastructure/persitence/seed"
)

var PersonaSeedCmd = &cobra.Command{
	Use:     "db:seed:persona",
	Aliases: []string{"m"},
	Short:   "Run Seeds for Personas",
}

func init() {
	PersonaSeedCmd.AddCommand(runSeedCmd)
}

var runSeedCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"m"},
	Short:   "Run seeds.",
	Run:     runSeed,
}

func runSeed(cmd *cobra.Command, args []string) {
	color.Blue(">>>Seed...  🚀 <<<")
	package_seed.ExecCommand()
	color.Blue(">>>Seed... 🚀 <<<")
}
