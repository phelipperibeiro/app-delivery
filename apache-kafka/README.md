# App Delivery (kafka)

## Descrição

Repositório do Apache Kafka (Backend)

## Configurar /etc/hosts

A comunicação entre as aplicações se dá de forma direta através da rede da máquina.
Para isto é necessário configurar um endereços que todos os containers Docker consigam acessar.

Acrescente no seu /etc/hosts:
```
127.0.0.1 host.docker.internal
```
Em todos os sistemas operacionais é necessário abrir o programa para editar o *hosts* como Administrator da máquina ou root.

## Rodar a aplicação

Execute os comandos na raiz do projeto:

```
mkdir -p es01/usr/share/elasticsearch/data
```

```
docker-compose up
```

Quando parar os containers do Kafka, lembre-se antes de rodar o **docker-compose up**, rodar o **docker-compose down** para limpar o armazenamento, senão lançará erro ao subir novamente.


## Observações

O Elasticsearch usa um **mmapfs** diretório por padrão para armazenar seus índices. É provável que os limites padrão do sistema operacional nas contagens de mmap sejam muito baixos, o que pode resultar em exceções de falta de memória.

No Linux, você pode aumentar os limites executando o seguinte comando como root:


```
sudo sysctl -w vm.max_map_count=262144
```

Para definir esse valor permanentemente, atualize a **vm.max_map_count** configuração em **/etc/sysctl.conf.** 
Para verificar após a reinicialização, execute:

```
sysctl vm.max_map_count
```


### Acessos

Acessar o confluent: 
```
http://localhost:9021/clusters.
```

Acessar o kibana:
```
http://localhost:5601.
```

Testar o elasticsearch:
```
curl -XGET http://127.0.0.1:9200
```
