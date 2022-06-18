package main

import (
	"os"
	"testing"
)

func TestInitEnvVariables(t *testing.T) {
	testApiPort := "test-port"
	os.Setenv("PORT", testApiPort)

	t.Run("Given all the env variables are set, no fatal log is thrown", func(t *testing.T) {
		initEnvVariables()
	})
}
