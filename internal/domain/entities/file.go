package entities

type File struct {
	ID     string
	Name   string
	Size   int64
	Status string // Estados: PENDING, UPLOADING, COMPLETED, etc.
}

/**
* NewFile é um "contrutor" que garante que todo
* arquivo nasça com o status correto
**/

func NewFile(name string, size int64) *File {
	return &File{
		Name:   name,
		Size:   size,
		Status: "PENDING",
	}
}
