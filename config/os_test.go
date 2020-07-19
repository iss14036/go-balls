package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	_ = os.Setenv("valid", "valid")
	expected := "valid"

	assert.Equal(t, expected, GetString("valid"))
}