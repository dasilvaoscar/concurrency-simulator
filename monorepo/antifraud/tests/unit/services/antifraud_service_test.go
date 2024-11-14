package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"concurrency-simulator/monorepo/antifraud/internal/services"
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

func TestValidateAntifraudWithHighRiskAmount(t *testing.T) {
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

func TestAntifraudDetectsHighRiskInstallments(t *testing.T) {
	expect := false

	sut := services.NewAntifraudAnalisysService()

	sutResponse, _ := sut.Execute(
		services.AntifraudAnalisysServiceData{
			FirstName: "Oscar",
			LastName: "da Silva",
			Amount: 100000.00,
			Installments: 48,
			Type: "credit",
			Status: "pending",
		},
	)

	assert.Equal(t, sutResponse, expect)
}

func TestAntifraudDetectsSuspiciousName(t *testing.T) {
	expect := false

	sut := services.NewAntifraudAnalisysService()

	sutResponse, _ := sut.Execute(
		services.AntifraudAnalisysServiceData{
			FirstName: "O",
			LastName: "da Silva",
			Amount: 50.00,
			Installments: 4,
			Type: "credit",
			Status: "pending",
		},
	)

	assert.Equal(t, sutResponse, expect)
}

func TestAntifraudDetectsSuspiciousLastName(t *testing.T) {
	expect := false

	sut := services.NewAntifraudAnalisysService()

	sutResponse, _ := sut.Execute(
		services.AntifraudAnalisysServiceData{
			FirstName: "Oscar",
			LastName: "D",
			Amount: 50.00,
			Installments: 4,
			Type: "credit",
			Status: "pending",
		},
	)
	
	assert.Equal(t, sutResponse, expect)
}

func TestAntifraudDetectsSuspiciousFullName(t *testing.T) {
	expect := false

	sut := services.NewAntifraudAnalisysService()

	sutResponse, _ := sut.Execute(
		services.AntifraudAnalisysServiceData{
			FirstName: "O",
			LastName: "D",
			Amount: 50.00,
			Installments: 4,
			Type: "credit",
			Status: "pending",
		},
	)
	
	assert.Equal(t, sutResponse, expect)
}