# Concurrency Simulator

## Descrição 

Um monorepo com um sistema de pagamentos simulado, com o objetivo de treinar alguns conceitos importantes e ferramentas que eu acho interessantes de se testar

## Requisitos

- [Funcionais](./docs/requisitos/rf.md)
- [Não funcionais](./docs/requisitos/rnf.md)

### System Flow Diagram

```mermaid
architecture-beta
    group api(cloud)[API Layer]
    group messaging(server)[Messaging Layer]
    group services(database)[Services Layer]
    group databases(disk)[Database Layer]
    
    service client(internet)[Cliente] in api
    service core_api(server)[Core API] in api
    
    service payment_topic(disk)[payment_topic] in messaging
    service push_notification_queue(disk)[push_notification_queue] in messaging
    service transaction_queue(disk)[transaction_queue] in messaging
    
    service customer_svc(database)[customer-svc] in services
    service antifraud_svc(database)[antifraud-svc] in services
    service transaction_svc(database)[transaction-svc] in services
    service notification_svc(database)[push-notification-svc] in services
    
    service dynamodb(database)[DynamoDB] in databases
    service postgres(database)[PostgreSQL] in databases
    
    client:R --> L:core_api
    core_api:R --> L:payment_topic
    payment_topic:R --> L:customer_svc
    payment_topic:R --> L:antifraud_svc
    customer_svc:R --> L:push_notification_queue
    antifraud_svc:R --> L:push_notification_queue
    antifraud_svc:R --> L:transaction_queue
    transaction_queue:R --> L:transaction_svc
    transaction_svc:R --> L:push_notification_queue
    push_notification_queue:R --> L:notification_svc
    
    customer_svc:B --> T:dynamodb
    notification_svc:B --> T:dynamodb
    antifraud_svc:B --> T:postgres
    transaction_svc:B --> T:postgres
```

## Ambientes Docker

### Iniciar containers

```zsh
docker compose up
```

## Utils

- [Reset de tópico no Kafka](./docs/kafka/como_deletar_topico.md)
- [Teste de envio de pagamento](./docs/http/POST.http)
