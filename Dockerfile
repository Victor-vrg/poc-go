# Etapa de build
FROM golang:1.23-alpine AS builder

# Definir o diretório de trabalho no container
WORKDIR /app

# Copiar os arquivos go.mod e go.sum primeiro, para aproveitar o cache de dependências
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar o restante dos arquivos do projeto
COPY . .

# Compilar o binário a partir do arquivo main.go localizado em cmd/
RUN go build -o /app/main ./cmd/main.go

# Etapa final: criar uma imagem mais leve
FROM alpine:latest

# Definir o diretório de trabalho na nova imagem
WORKDIR /app

# Copiar o binário da etapa de build
COPY --from=builder /app/main .

# Expor a porta da aplicação (ajuste conforme necessário)
EXPOSE 8080

# Comando de inicialização
CMD ["/app/main"]
