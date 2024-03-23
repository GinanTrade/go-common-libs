package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigInit(t *testing.T) {
	// Config.Load()
	wd, _ := os.Getwd()
	t.Log("Testing config init", wd)
	assert.NotEqual(t, Config.Get("MONGODB_URL"), "")
}
