package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	assert.NotNil(t, AppConfig())
}