package ports

import "dlss/internal/domain/entities"

/**
* StorageRepository define o contrato para persistência
* física de arquivos.
* Esta interface permite que o domínio salve dados sem saber "onde" e "como".
**/
type StorageRepository interface {
	/**
	* Save recebe a entidade File e os bytes do arquivo
	* para persistir no disco
	**/
	Save(file entities.File, data []byte) error
}
