package filesystem

import (
	"dlss/internal/domain/entities"
	"os"
	"path/filepath"
)

/**
* FileSystemAdapter implementa a interface StorageRepository salvando em disco.
**/
type FilesSystemAdapter struct {
	basePath string
}

/**
* NewFileSystemAdapter define onde os arquivos serão
* salvos
**/
func NewFileSystemAdapter(basePath string) *FilesSystemAdapter {
	return &FilesSystemAdapter{basePath: basePath}

}

/**
* Save grave os bytes do arquivo no caminho físico especificado
**/
func (f *FilesSystemAdapter) Save(file entities.File, data []byte) error {
	/**
	* 0 - indicando que deve s er interpretado como octal
	* 7 Dono - Leitura, escrita e execução
	* 5 Grupo - Leitura e execução
	* 5 Outros - Leitura e execução
	*
	* Garante que quando o usuário fizer o primeiro upload,
	* o GO consiga criar a pasta no disco e tenha permissão para
	* colocar arquivos lá dentro
	**/
	if err := os.MkdirAll(f.basePath, 0755); err != nil {
		return err
	}

	fullPath := filepath.Join(f.basePath, file.Name)

	/**
	* 0 - Octal
	* 6 Dono - Leitura e escrita
	* 4 Grupo - Leitura apenas
	* 4 Outros - Leitura apenas
	*
	* Garante que os arquivos salvos no disco estejam disponíveis para leitura
	* posterior, mas protegidos contra escrita de outros usuários que não sejam
	* o dono do processo do backend.
	**/
	return os.WriteFile(fullPath, data, 0644)
}
