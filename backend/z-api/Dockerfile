   # Use uma imagem base do Go
   FROM golang:1.20

   # Defina o diretório de trabalho
   WORKDIR /app

   # Copie os arquivos go.mod e go.sum e instale as dependências
   COPY go.mod go.sum ./
   RUN go mod download

   # Copie o restante do código
   COPY . .

   # Compile o aplicativo
   RUN go build -o main .

   # Comando para rodar o aplicativo
   CMD ["./main"]