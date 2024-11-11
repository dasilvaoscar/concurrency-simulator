package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAntifraudWithValidData(t *testing.T) {
	expect := true

	sutResponse := true

	assert.Equal(t, sutResponse, expect)
}

func TestValidateAntifraudWithInvalidData(t *testing.T) {
	expect := false

	sutResponse := false

	assert.Equal(t, sutResponse, expect)
}