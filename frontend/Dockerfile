# Building frontend image
FROM node:23-alpine3.20 AS builder

WORKDIR /app

COPY package.json .
COPY package-lock.json .

RUN npm install

COPY . .

RUN npm run build

# Etapa 2: Servir com NGINX
FROM nginx:alpine

# Remove a configuração default do NGINX
RUN rm -rf /usr/share/nginx/html/*

# Copia os arquivos buildados da etapa anterior
COPY --from=builder /app/dist /usr/share/nginx/html

# Expõe a porta padrão do NGINX
EXPOSE 80

# Comando padrão
CMD ["nginx", "-g", "daemon off;"]