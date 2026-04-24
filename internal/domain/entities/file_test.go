package entities

import "testing"

func TestNewFileStatus(t *testing.T) {
	file := NewFile("projeto.zip", 1024)

	if file.Status != "PENDING" {
		t.Errorf("Esperando status PENDING, recebido %s", file.Status)
	}

}
