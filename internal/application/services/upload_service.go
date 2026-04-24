package services

import (
	"dlss/internal/application/ports"
	"dlss/internal/domain/entities"
	"dlss/internal/domain/services"
)

/**
* UploadService orquestra o fluxo de entrada de arquivos no sistema.
**/
type UploadService struct {
	storage ports.StorageRepository
}

/**
* NewUploadService é o construtor do serviço de aplicação.
**/
func NewUploadService(s ports.StorageRepository) *UploadService {
	return &UploadService{storage: s}
}

/**
*ExecuteUpload coordena a validação de quota e o salvamento físico.
**/
func (u *UploadService) ExecuteUpload(file entities.File, data []byte, usedSpace int64) error {
	err := services.ValidateQuota(&file, usedSpace)

	if err != nil {
		return err
	}

	return u.storage.Save(file, data)
}
