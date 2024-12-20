# Stage de build
FROM golang:1.23.2-alpine AS builder

WORKDIR /app

# Instalar dependências de build e protoc
RUN apk add --no-cache gcc musl-dev protobuf protobuf-dev

# Instalar protoc-gen-go e protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    && export PATH="$PATH:$(go env GOPATH)/bin"

# Copiar go.mod e go.sum
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod tidy

# Copiar código fonte e proto
COPY . .

# Compilar aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o main .

# Stage final
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]