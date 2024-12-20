# Etapa de build (builder)
FROM golang:1.23.2 as builder

# Defina o diretório de trabalho no container
WORKDIR /app

# Instale o protobuf compiler (protoc) e os plugins Go
RUN apt-get update && apt-get install -y \
    unzip \
    protobuf-compiler

# Instale os plugins protoc-gen-go e protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    && export PATH="$PATH:$(go env GOPATH)/bin"

# Copie o código-fonte
COPY . .

# Baixe as dependências
RUN go mod tidy

# Compile os arquivos proto
#RUN protoc --go_out=. --go-grpc_out=. proto/*.proto

# Compile o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/meeting ./cmd

# Etapa de execução (runtime)
FROM alpine:latest

# Criação de um volume para persistência
VOLUME /app/data

# Copie o binário do estágio anterior
COPY --from=builder /app/bin/meeting /bin/meeting

# Defina o entrypoint para executar o binário
ENTRYPOINT ["/bin/meeting"]
