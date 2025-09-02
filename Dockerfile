# Etapa de build
FROM golang:1.25-alpine AS builder

WORKDIR /builder

# Copiar e baixar dependências primeiro
COPY go.mod go.sum ./
RUN go mod download

# Copiar todo o código
COPY . .

# Compilar binário (garantindo path correto)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Imagem final mínima
FROM scratch

WORKDIR /app
COPY --from=builder /builder/main .
COPY --from=builder /builder/.env .

EXPOSE 8080

CMD ["/app/main"]