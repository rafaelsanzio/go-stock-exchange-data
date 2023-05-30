# Imagem base
FROM golang:1.19-alpine

# Instalando as dependências
RUN apk update && apk add git

# Copiando os arquivos do projeto
COPY . /

# Diretório de trabalho
WORKDIR /

# Compilando a aplicação
RUN go mod download
RUN go build ./main.go

# Expondo a porta 9090
EXPOSE 9091

# Comando para iniciar a aplicação
CMD ["./main"]
