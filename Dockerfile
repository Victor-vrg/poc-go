# Etapa de build
FROM golang:1.18-alpine AS builder

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos de dependência
COPY go.mod ./
COPY go.sum ./

# Baixa as dependências
RUN go mod download

# Copia o restante do código para o container
COPY . .

# Compila a aplicação
RUN go build -o /app/main .

# Etapa final: cria uma imagem leve para rodar o binário
FROM alpine:latest

WORKDIR /root/

# Copia o binário compilado da etapa anterior
COPY --from=builder /app/main .

# Define o comando padrão ao iniciar o container
CMD ["./main"]
