package main

import (
	"github.com/juanfer2/persona_5/cmd"
	"github.com/juanfer2/persona_5/src/shared/infrastructure/config"
)

func main() {
	config.LoadEnv()
	cmd.Execute()
}
