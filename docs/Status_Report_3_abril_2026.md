Status Atual: Integração Core-Storage finalizada

|Componente | Localização | Responsabilidade |
| :--- | :--- | :--- |
| Entidade | `domain/entities/file.go` | Define o arquivo e status PENDING
| Serviço de Domínio | `domain/services/quota_service.go` | Garante a regra SaaS de 10MB|
| Port (Interface) | `application/ports/storage_repository.go` | Contrato de salvamento| 
| Serviço de Aplicação | `application/services/upload_service.go` | Maestro do upload | 
Adapter de saída | `adapters/secondary/filesystem/...` | Implementação real de escrita no disco