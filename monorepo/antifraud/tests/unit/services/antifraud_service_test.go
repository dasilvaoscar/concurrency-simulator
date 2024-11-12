package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"monorepo/antifraud/internal/services"
)

func TestValidateAntifraudWithValidData(t *testing.T) {
	expect := true

	sut := services.NewAntifraudAnalisysService()

	sutResponse, _ := sut.Execute(
		services.AntifraudAnalisysServiceData{
			FirstName: "Oscar",
			LastName: "da Silva",
			Amount: 50.00,
			Installments: 4,
			Type: "credit",
			Status: "pending",
		},
	)

	assert.Equal(t, sutResponse, expect)
}

func TestValidateAntifraudWithInvalidData(t *testing.T) {
	expect := false

	sut := services.NewAntifraudAnalisysService()

	sutResponse, _ := sut.Execute(
		services.AntifraudAnalisysServiceData{
			FirstName: "Oscar",
			LastName: "da Silva",
			Amount: 1000000.00,
			Installments: 4,
			Type: "credit",
			Status: "pending",
		},
	)

	assert.Equal(t, sutResponse, expect)
}