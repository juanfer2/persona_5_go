package utilities

import (
	"os"
	"testing"

	"github.com/juanfer2/persona_5/src/shared/utilities"
	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("MODE", "TEST")
	assert := assert.New(t)
	value := utilities.GetEnv("MODE")

	assert.Equal(value, "TEST", "they should be truthy")
}
