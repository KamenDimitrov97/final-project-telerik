package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDefaultConfig(t *testing.T) {
	config, err := Get()
	assert.Nil(t, err)

	assert.Equal(t, ":3333", config.BindAddr)
}
