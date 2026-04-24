package services

import (
	"dlss/internal/domain/entities"
	"testing"
)

type MockStorage struct {
	SaveCalled bool
}

func (m *MockStorage) Save(file entities.File, data []byte) error {
	m.SaveCalled = true
	return nil
}

func TestExecuteUpload_Success(t *testing.T) {
	mock := &MockStorage{}
	service := NewUploadService(mock)

	file := entities.NewFile("test.txt", 1024)

	err := service.ExecuteUpload(*file, []byte("conteudo"), 0)

	if err != nil {
		t.Errorf("Não esperava erro, mas obteve: %v", err)
	}

	if !mock.SaveCalled {
		t.Error("O StorageRepository deveria ter sido chamado")
	}
}
