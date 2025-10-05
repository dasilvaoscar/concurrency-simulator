# Concurrency Simulator

## Descrição 

Um monorepo com um sistema de pagamentos simulado, com o objetivo de treinar alguns conceitos importantes e ferramentas que eu acho interessantes de se testar

## Requisitos

- [Funcionais](./docs/requisitos/rf.md)
- [Não funcionais](./docs/requisitos/rnf.md)

## System Flow Diagram

```mermaid
flowchart LR
    Client[Cliente] --> CoreAPI[Core API]
    CoreAPI --> PaymentTopic[payment_topic]
    
    PaymentTopic --> CustomerSvc[customer-svc]
    PaymentTopic --> AntifraudSvc[antifraud-svc]
    
    CustomerSvc --> PushNotificationQueue[push_notification_queue]
    AntifraudSvc --> PushNotificationQueue
    AntifraudSvc --> TransactionQueue[transaction_queue]
    
    TransactionQueue --> TransactionSvc[transaction-svc]
    TransactionSvc --> PushNotificationQueue
    PushNotificationQueue --> NotificationSvc[push-notification-svc]
    
    CustomerSvc --> DynamoDB[(DynamoDB)]
    NotificationSvc --> DynamoDB
    AntifraudSvc --> Postgres[(PostgreSQL)]
    TransactionSvc --> Postgres
    
    style Client fill:#e1f5fe,color:#000000
    style CoreAPI fill:#f3e5f5,color:#000000
    style PaymentTopic fill:#fff3e0,color:#000000
    style CustomerSvc fill:#e8f5e8,color:#000000
    style AntifraudSvc fill:#e8f5e8,color:#000000
    style TransactionSvc fill:#e8f5e8,color:#000000
    style NotificationSvc fill:#e8f5e8,color:#000000
    style PushNotificationQueue fill:#fff3e0,color:#000000
    style TransactionQueue fill:#fff3e0,color:#000000
    style DynamoDB fill:#ffebee,color:#000000
    style Postgres fill:#ffebee,color:#000000
```

## Ambientes Docker

### Iniciar containers

```zsh
docker compose up
```

## Utils

- [Reset de tópico no Kafka](./docs/kafka/como_deletar_topico.md)
- [Teste de envio de pagamento](./docs/http/POST.http)
