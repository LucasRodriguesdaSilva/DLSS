# Status Report: DLSS

Data: Abril de 2026
Metodologia: TDD + Arquitetura Hexagonal 
Ambiente: Híbrido (Desenvolvimento no Windows / infraestrutura no WSL2)

## 1. Estrutura de Diretórios Atual 

O projeto foi inicializado utilizando o framework Wails e organizado para isolar o domínio:

- docker/: Contém uma cópia do docker-compose.yml que é utilizado no WSL2 para serviçoes de infraestrutura (Redis e Storage)
- /internal/domain/entities/: Local das entidades puras e seus testes de unidade.
- /internal/domain/services/: (Em desenvolvimento) Lógica de decisão que envolve múltiplas entidades.
- /internal/adapters/: (Planjeado) Implementações técnicas de saída (SQLite, Redis, FileSystem).

## 2. Componentes Implementados 

### Infraestrutura (Simulação de Servidor)

- Serviço: Redis (Porta 6379) para gestão de chunks e upload resumbale.
- Serviço: Storage Node (Volume Docker) para simular o "Disco" de armazenamento massico

### Domínio: Entidade **File**
Local /internal/domain/entities/file.go

- Struct `File`: Armazena metadados como `ID`, `Name`, `Size` e `Status`.
- Função `NewFile(name, size)`: Construtor que garante a regra de negócio de que todo arquivo nasce com o estado `PENDING`.

## 3. Cobertura de Testes (TDD)

| Teste | Objetivo | Status |
| :--- | :--- | :--- |
| TestNewFileStatus |Garante que novos arquivos iniciem como PENDING. | PASS 

## 4. O que está em desenvolvimento (Próximos Passos)
1. Validação de Quota (Service): Implementar o `QuotaService` para validar se um upload respeita o limite de 10MB do plano FREE.
2. Repositórios (Ports): Definir as interfaces (ports) para que o domíniuo consiga salvar arquivos e consultar o SQLite sem depender da implementação real.
3. Máquina de Estados: Evoluir a entidade `FILE` para suportar as transições de estado: `UPLOADING`, `COMPLETED`, `FAILED`, etc.

## 5. Pendências (Roadmap Futuro)

- [ ] Implementar gRPC para streaming binário de arquivos.
- [ ] Criar interface gráfica em Vue3 + Tailwind via Wails.
- [ ] Configurar Ghost Detection (detectar arquivos apagados manualmente no disco).
- [ ] Simular sistema de pagamentos sandbox para upgrade de quota.