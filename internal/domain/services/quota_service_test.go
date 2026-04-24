package services

import (
	"dlss/internal/domain/entities"
	"testing"
)

func TestValidateQuota(t *testing.T) {
	t.Run("Deve retornar erro se exceder os 10MB do plano FREE", func(t *testing.T) {
		// 11 MB
		file := entities.NewFile("grande.zip", (11 * 1024 * 1024))
		usedSpace := int64(0)

		err := ValidateQuota(file, usedSpace)

		if err == nil {
			t.Error("Esperado erro de quota excedida, mas recebeu nil")
		}
	})

	t.Run("Deve permitir se estiver dentro do limite", func(t *testing.T) {
		// 5MB
		file := entities.NewFile("pequeno.zip", (5 * 1024 * 1024))
		usedSpace := int64(0)

		err := ValidateQuota(file, usedSpace)

		if err != nil {
			t.Errorf("Não esperado erro, mas recebeu: %v", err)
		}
	})
}
