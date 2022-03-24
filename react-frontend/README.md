# App Delivery (Aplicação/front-end)

## Descrição

Repositório do front-end feito com React.js (Front-end)

**Importante**: A aplicação do Apache Kafka, Golang e Nest.js deve estar rodando primeiro.

## Configurar /etc/hosts

A comunicação entre as aplicações se dá de forma direta através da rede da máquina.
Para isto é necessário configurar um endereços que todos os containers Docker consigam acessar.

Acrescente no seu /etc/hosts:
```
127.0.0.1 host.docker.internal
```
Em todos os sistemas operacionais é necessário abrir o programa para editar o *hosts* como Administrator da máquina ou root.

## Serviços da Google que precisam ser habilitados

- Maps JavaScript API
- Directions API

## Rodar a aplicação

Execute os comandos:

```
docker-compose up
```

Acessar http://localhost:3001.