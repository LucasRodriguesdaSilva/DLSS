package filesystem

import (
	"dlss/internal/domain/entities"
	"os"
	"path/filepath"
	"testing"
)

func TestFileSystemSave(t *testing.T) {
	tmpDir := t.TempDir()
	adapter := NewFileSystemAdapter(tmpDir)
	file := entities.NewFile("integracao.txt", 10)
	content := []byte("hello dlss")

	err := adapter.Save(*file, content)

	if err != nil {
		t.Fatalf("Erro ao salvar arquivo: %v", err)
	}

	expectedPath := filepath.Join(tmpDir, "integracao.txt")

	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Error("O arquivo físico não foi criado no diretório esperando")
	}
}
