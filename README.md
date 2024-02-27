# Executando a Aplicação

Este documento fornece instruções passo a passo para executar a aplicação.

## Pré-requisitos

- Go (versão 1.16 ou superior)
- Docker e Docker Compose (para o banco de dados MySQL e RabbitMQ)

## Passos para Execução

1. Clone o repositório para sua máquina local.
2. Navegue até a pasta do projeto.
3. Inicie os serviços do Docker (MySQL e RabbitMQ) com o seguinte comando:

```sh
$ docker-compose up -d
```

4. Instale as dependências do projeto:

```sh
$ go mod tidy
```

5. Navegue até a pasta cmd/ordersystem e execute o arquivo main.go:

```sh
$ go run main.go
```

## Serviços e Portas

A aplicação inicia vários serviços que escutam nas seguintes portas:

- Servidor Web: Porta 8080 (http://localhost:8080)
- Servidor gRPC: Porta 50051
- Servidor GraphQL: Porta 8081 (http://localhost:8081/query)

## Testando a Aplicação

Você pode testar a aplicação usando os arquivos .http na pasta api. Estes arquivos contêm exemplos de requisições HTTP que podem ser enviadas para a aplicação. Você pode usar uma extensão como a REST Client no Visual Studio Code para enviar essas requisições diretamente do editor.