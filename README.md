Distributed Local Storage System (DLSS)
Distributed Local Storage System (DLSS) é um sistema de Object Storage local que simula um serviço SaaS (Google Drive/S3), utilizando streaming via gRPC, controle de quota, upload resumable e armazenamento distribuído em disco/volume Docker.

📌 Visão Geral
O DLSS é um sistema de gerenciamento de arquivos de alto desempenho projetado para simular um serviço de cloud storage, porém executando localmente.
A arquitetura separa:
Metadados e aplicação no disco principal (C:)
Armazenamento massivo de arquivos em um volume Docker
O sistema foi pensado para aprendizado profundo em tecnologias utilizadas por Big Techs, com foco em performance, streaming binário, concorrência, persistência, arquitetura limpa e testes automatizados.

🎯 Objetivos do Projeto
Implementar um Object Storage local com comportamento SaaS
Exercitar Go (Golang) para sistemas de infraestrutura e I/O
Utilizar gRPC streaming para upload eficiente
Aplicar Arquitetura Hexagonal (Ports & Adapters)
Utilizar TDD como metodologia principal
Suportar quota por usuário e capacidade global
Suportar planos de assinatura e pagamentos sandbox
Ter um painel administrativo para monitoramento e expansão do sistema

🧱 Stack Tecnológica
Backend / Core
Go (Golang)
gRPC + Protocol Buffers
SQLite (metadados)
Redis (chunks temporários)
Docker Volume (armazenamento físico)
Desktop UI
Wails
Vue 3 + TypeScript + Tailwind
Arquitetura e Qualidade
Arquitetura Hexagonal
TDD (Test Driven Development)
Logs estruturados
Observabilidade básica

🧠 Conceitos Principais
Separação Física (Simulação SaaS)
Disco C: Metadados + aplicação (SQLite, config, usuários)
Disco D: Armazenamento massivo de arquivos (Docker Volume)
Modelo SaaS Simulado
Usuários possuem planos (FREE, PRO, etc.)
Cada plano define uma quota individual
O sistema possui uma capacidade global total
Admin pode “comprar” mais capacidade global para evitar bloqueios

📂 Estrutura do Projeto (Sugestão)
dlss/
├── cmd/
│   ├── server/               # Entry point do backend
│   └── desktop/              # Entry point Wails
│
├── internal/
│   ├── domain/               # Regras de negócio puras
│   │   ├── entities/
│   │   ├── services/
│   │   └── errors/
│   │
│   ├── application/          # Use cases / orchestration
│   │   ├── ports/
│   │   └── usecases/
│   │
│   ├── adapters/
│   │   ├── primary/          # Entradas (gRPC, Wails)
│   │   └── secondary/        # Saídas (SQLite, Redis, FS)
│   │
│   ├── infra/
│   │   ├── database/
│   │   ├── redis/
│   │   └── filesystem/
│   │
│   └── config/
│
├── proto/                    # Protobuf definitions
├── docker/                   # Docker compose / volumes
├── scripts/                  # scripts auxiliares
├── docs/                     # documentação extra
└── README.md


🏛️ Arquitetura (Hexagonal)
O sistema segue Arquitetura Hexagonal, garantindo que o domínio não dependa de frameworks.
Camadas
Domain (Core)
Responsável por:
entidades (User, File, Plan, Payment)
validação de quota
regras de estado de arquivo
erros de domínio
Application (Use Cases)
Responsável por:
orquestrar upload e persistência
gerenciar transações de plano/pagamento
coordenar Redis + SQLite + filesystem
Adapters (Ports & Adapters)
Primary Adapters (Input):
gRPC Server
Wails bridge (UI Desktop)
Secondary Adapters (Output):
SQLite repository
Redis chunk storage
filesystem driver (Docker Volume)

🔁 Fluxos Principais
Upload de Arquivo (gRPC Streaming)
Usuário solicita upload
Backend valida:
quota individual do usuário
capacidade global do sistema
Arquivo entra como PENDING
Upload inicia → estado UPLOADING
Chunks são enviados via gRPC stream
Redis armazena progresso temporário (resumable)
Backend grava dados no volume
Hash é validado
Status vira COMPLETED

Upgrade de Plano (Sandbox Payment)
Usuário escolhe plano maior
Sistema cria pagamento PENDING
API sandbox confirma → PAID
Sistema atualiza plano e quota do usuário
Evento é registrado no histórico

Expansão de Capacidade Global (Admin)
Admin monitora uso total
Admin “compra” mais espaço
Capacidade global aumenta
Uploads voltam a funcionar normalmente

📌 Máquina de Estados do Arquivo
Estado
Descrição
PENDING
arquivo enfileirado aguardando upload
UPLOADING
upload em andamento
PAUSED
upload interrompido e retomável
COMPLETED
upload finalizado com sucesso
FAILED
falha irreversível no upload
CONFLICT
conflito de nome detectado


📏 Regras de Quota
Quota Individual
Antes de iniciar upload:
used_space + file_size <= user_quota

Se exceder → upload bloqueado imediatamente.
Quota Global
Além da quota individual, existe uma capacidade global:
global_used + file_size <= global_capacity

Se exceder → upload bloqueado para todos os usuários.

💳 Planos e Pagamentos (Sandbox)
O sistema simula um modelo SaaS com upgrade.
Status de Pagamento
PENDING
PAID
FAILED
Upgrade só ocorre após confirmação (PAID).

🗃️ Persistência
SQLite (Fonte da Verdade)
Responsável por:
usuários
planos
arquivos (metadados)
status de upload
pagamentos e histórico
quota global e configurações
Redis (Estado Volátil)
Responsável por:
chunks temporários para resumable upload
tracking de progresso por file_id
Docker Volume (Storage Massivo)
Responsável por:
armazenar os arquivos físicos do sistema

🧾 Modelo de Dados (High Level)
users
id
name
email
password_hash
plan_id
used_space_bytes
created_at
plans
id
name
quota_bytes
price
active
files
id
user_id
name
size_bytes
hash
status
created_at
updated_at
payments
id
user_id
plan_id
status
provider
created_at
system_storage
id
total_capacity_bytes
used_capacity_bytes
updated_at

⚠️ Edge Cases Tratados
Disco/Volume indisponível → upload entra em PAUSED
Queda de energia → uploads UPLOADING voltam para PAUSED no boot
Arquivo apagado manualmente no volume → marcado como MISSING (Ghost Detection)
Downgrade de plano com usuário acima da quota → bloquear novos uploads
Pagamento confirmado mas falha no upgrade → reprocessamento idempotente

🔐 Segurança (Regras Iniciais)
Usuário só acessa seus próprios arquivos
Autenticação obrigatória
Auditoria de ações críticas:
upload
delete
upgrade
mudanças de quota
expansão global
Nota: como é um sistema local, o foco inicial é consistência e controle. Criptografia pode ser adicionada em versões futuras.

📊 Observabilidade
O sistema deve fornecer:
logs estruturados por request/upload
métricas básicas:
throughput de upload
taxa de falhas
tempo médio de upload
espaço usado individual e global

🧪 Testes (TDD)
Tipos de Teste
Unitários (domínio puro)
Integração (SQLite/Redis/FileSystem)
Testes de stream gRPC (simulação de upload grande)
Estratégia
Domínio deve ser testável sem dependências externas
Adapters devem ter testes de integração com containers

🚀 Como Rodar Localmente
Pré-requisitos
Go instalado
Docker + Docker Compose
Node.js (para Wails frontend)
Redis container (via compose)
Subir Infra (Redis + Volume)
docker compose up -d

Rodar Backend
go run ./cmd/server

Rodar Desktop (Wails)
wails dev


📦 Roadmap
V1 — MVP Core
upload via gRPC streaming
SQLite metadados
Docker Volume storage
quota FREE fixa (10MB)
UI básica de upload
V2 — Resumable Upload
Redis chunks
crash recovery
background polling / ghost detection
V3 — Multiusuário + Planos
cadastro/login real
planos configuráveis
quota por usuário
V4 — Pagamentos Sandbox
simulação de pagamentos
upgrades automáticos
V5 — Admin Console
monitoramento global
compra/expansão de capacidade global

📌 Considerações Futuras (Nice to Have)
criptografia AES por arquivo
compressão automática
versionamento de arquivos
compartilhamento entre usuários
replicação real em rede (multi-node)
integração com gateways reais (Stripe, MercadoPago)

📜 Licença
Este projeto é desenvolvido para fins educacionais e experimentais no contexto de "Build to Learn".

