## Entrar na linha de comando
docker-compose exec kafka1 bash

## Deletar tópico
kafka-topics --bootstrap-server localhost:19092 --delete --topic payment_topic

## Recriar
kafka-topics --bootstrap-server localhost:19092 --create --topic payment_topic --partitions 1 --replication-factor 1