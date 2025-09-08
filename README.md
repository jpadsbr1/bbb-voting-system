
# 🗳️ BBB Voting API  

API que simula a votação do **Big Brother Brasil**, projetada para ser **escalável, resiliente e performática**, capaz de lidar com milhares de votos por segundo. Feita com **[Golang](https://go.dev/)**.

---

## 📐 Arquitetura

A arquitetura segue princípios de **Clean Architecture** e **Event-Driven Architecture**, garantindo **baixo acoplamento** e **alta escalabilidade**:

- **API:**
  - Recebe votos via REST (`POST /vote`).
- **Redis (Cache/Counter):**
  - Mantém contadores de votos em tempo real.
- **Postgres (Storage):**
  - Persistência definitiva (participantes, eliminações, histórico de votos por hora).

---

## 🛠️ Ferramentas Utilizadas

- **[Go](https://go.dev/)** → linguagem principal da API/Worker.
- **[Postgres](https://www.postgresql.org/)** → banco relacional para persistência duradoura.  
- **[Redis](https://redis.io/)** → cache em memória para contagem rápida.
- **[Docker & Docker Compose](https://www.docker.com/)** → orquestração local de toda a stack.  
- **[Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)** → organização em camadas (Domain, UseCases, Interfaces, Infra).  

---

## 📂 Estrutura de Pastas

```
/cmd
  /app/main.go        # binário único
/internal
  /config             # ferramentas de configuração
  /delivery
    /http             # handlers de configuração de rotas
  /domain             # entidades puras (Vote, Participant, BigWall)
  /infrastructure     # implementações concretas (Redis, Postgres)
  /repository         # interações com infra (Banco de dados e cache)
  /usecase            # casos de uso (Serviços de voto, participante e paredão)
```

---

## ⚡ Fluxo de Votação

1. Usuário envia voto:
   ```http
   POST /vote
   {
     "participant_id": x
     "bigwall_id": x
   }
   ```
2. API recebe voto e incrementa uma contagem no Redis, uma para totais, outra para voto periódico por hora.
5. Periodicamente (5 segundos), o Worker grava os totais no Postgres, em tabelas totais, por participante e por hora.

---

## 🧨 Rotas

Adicionar participante
```bash
curl --request POST \
  --url http://localhost:8080/participant \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/11.5.0' \
  --data '{
	"name": x
}'
```
Retornar todos os participantes
```bash
curl --request GET \
  --url http://localhost:8080/participants \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/11.5.0'
```
Criar paredão
```bash
curl --request POST \
  --url http://localhost:8080/bigwall/create \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/11.5.0' \
  --data '{
	"participant_ids": [x, x]
}'
```
Retornar participantes que estão no paredão
```bash
curl --request GET \
  --url http://localhost:8080/bigwall/participants/:BigWallID \
  --header 'User-Agent: insomnia/11.5.0'
```
Retornar paredão atual
```bash
curl --request GET \
  --url http://localhost:8080/bigwall \
  --header 'User-Agent: insomnia/11.5.0'
```
Finalizar paredão e eliminar participante
```bash
curl --request PATCH \
  --url http://localhost:8080/bigwall/end/:BigWallID \
  --header 'User-Agent: insomnia/11.5.0'
```
Votar
```bash
curl --request POST \
  --url http://localhost:8080/vote \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/11.3.0' \
  --data '{
	"bigwall_id" : x,
	"participant_id": x
}'
```
Retornar total de votos do paredão
```bash
curl --request GET \
  --url http://localhost:8080/votes/total/:BigWallID \
  --header 'User-Agent: insomnia/11.3.0'
```
Retornar votos por participante do paredão
```bash
curl --request GET \
  --url http://localhost:8080/votes/participant \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/11.5.0' \
  --data '{
	"participant_id": x,
	"bigwall_id": x
}'
```
Retornar votos por hora do paredão
```bash
curl --request GET \
  --url http://localhost:8080/votes/hourly/:BigWallID \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/11.5.0'
```
## ▶️ Como Rodar

#### Requerimentos
* Golang 1.25
* Docker e Docker Compose
* Make e Golang-Migrate

### 1. Clone o repositório
```bash
git clone https://github.com/seu-usuario/bbb-voting-api.git
cd bbb-voting-api
```

### 2. Suba os serviços
```bash
docker-compose up --build
```

### 3. Rodar migrations de Postgres
```bash
make migrate-up
```

Isso vai iniciar:
- **API** → `localhost:8080`
- **Postgres** → `localhost:5432`  
- **Redis** → `localhost:6379`

### 3. Adicionar participantes
```bash
curl --request POST \
  --url http://localhost:8080/participant \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/11.5.0' \
  --data '{
	"name": x
}'
```
### 4. Criar um paredão
```bash
curl --request POST \
  --url http://localhost:8080/bigwall/create \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/11.5.0' \
  --data '{
	"participant_ids": [x, x]
}'
```
### 5. Enviar um voto
```bash
curl --request POST \
  --url http://localhost:8080/vote \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/11.3.0' \
  --data '{
	"bigwall_id" : x,
	"participant_id": x
}'
```
### 6. Finalizar Paredão
```bash
curl --request PATCH \
  --url http://localhost:8080/bigwall/end/:BigWallID \
  --header 'User-Agent: insomnia/11.5.0'
```

---

## ✅ Próximos Passos

- Implementar Kafka para enfileiramento de votos, aumentando escalabilidade
- Adicionar testes unitários (com mocks de ports).
- Adicionar métricas Prometheus + Grafana.
- Adicionar ELK (Elastic + Logstash + Kibana) para logs.
- Implementar CI/CD com GitHub Actions.
- Provisionar infraestrutura na AWS através de IaC com Terraform.
- Criar um front-end para interação e simulação real de um paredão.
