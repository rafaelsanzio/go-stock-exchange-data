# Stock Exchange Data 💲

O objetivo deste projeto é simular operações da bolsa de valores e enviá-las para outro sistema por meio de um servidor WebSocket. Utilizamos um scheduler para gerar e enviar dados de operações da bolsa em intervalos regulares. O servidor WebSocket recebe esses dados e os processa conforme necessário.

O projeto visa reproduzir um cenário realista de transmissão de informações da bolsa de valores em tempo real. O scheduler gera dados simulados de operações, como preços de ações, mudanças de valores e outros detalhes relevantes. Esses dados são enviados para o servidor WebSocket, que é responsável por receber as informações e executar o processamento adequado.

**Observação: Este projeto contém dados fictícios para operar a simulação.**

## Iniciando ▶️

### Pré-requisitos

Antes de começar, certifique-se de ter o seguinte instalado:

- Docker: https://www.docker.com/

### Instalação

1. Clone o repositório: https://github.com/rafaelsanzio/go-stock-exchange-data
2. Crie um arquivo `.env` com base no exemplo `.env.example` fornecido. Se necessário, ajuste os valores conforme sua preferência.
3. É necessário criar uma rede Docker para permitir a conexão com outros serviços:
   ```sh
   docker network create app_network
   ```
4. Por fim, para iniciar a aplicação, execute os comandos de build e up:
   ```sh
   docker-compose build && docker-compose up
   ```
