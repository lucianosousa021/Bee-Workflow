# Estágio de build
FROM node:20-alpine AS builder

WORKDIR /app

# Copiar arquivos de package.json
COPY package*.json ./

# Instalar dependências
RUN npm install

# Copiar o resto dos arquivos
COPY . .

# Criar build de produção
RUN npm run build

# Estágio de produção
FROM node:20-alpine AS runner

WORKDIR /app

# Copiar package.json e package-lock.json
COPY package*.json ./

# Instalar apenas dependências de produção
RUN npm install --production

# Copiar build e arquivos públicos do estágio anterior
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/next.config.ts ./

# Expor porta
EXPOSE 3000

# Iniciar aplicação
CMD ["npm", "start"] 