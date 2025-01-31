FROM golang:1.19-alpine

WORKDIR /app

# Copia apenas os arquivos necessários para baixar as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila o aplicativo
RUN go build -o users .

# Dá permissão de execução ao binário
RUN chmod +x users

EXPOSE 3000

CMD ["./users"]