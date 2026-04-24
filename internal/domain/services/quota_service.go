package services

import (
	"dlss/internal/domain/entities"
	"errors"
)

// Limite do plano FREE definido no PRD
const MaxFreeQuotaBytes = 10 * 1024 * 1024

// ValidateQuota verifica se o novo arquivo cabe na quota do usuário
func ValidateQuota(file *entities.File, usedSpace int64) error {
	if (usedSpace + file.Size) > MaxFreeQuotaBytes {
		return errors.New("quota exceed: seu plano FREE permite apenas 10MB")
	}
	return nil
}
