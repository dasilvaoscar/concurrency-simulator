# Harmonia Financeira

## Contexto

Inicialmente esse repositório foi criado para simular um sistema de pagamento com alta concorrência. Porém passei a usa-lo como meu projeto de [TCC](https://docs.google.com/document/d/1JjoOGEuK71dnkTWEYsoSsqkzM1umRoWZ/edit?usp=sharing&ouid=113161032001964584333&rtpof=true&sd=true)
da **pós-graduação de arquitetura**.

## Descrição 

Esse projeto é um monorepo com um sistema de pagamentos simulado e gestão de contas familiar em construção.

## Requisitos

- [Funcionais](./docs/requisitos/rf.md)
- [Não funcionais](./docs/requisitos/rnf.md)

## Domínios de negócio

```mermaid
---
config:
  theme: redux
  layout: dagre
  look: neo
---
flowchart TB
    n1(["Banking"]) --> n2["Pix"] & n3["Boleto"] & n5["TED"]
    A(["Gestão de gastos"]) --> n4["Dashboards - Indicadores de gastos exacerbados"] & n6["Alerta de gastos"] & n7["Unificação de controle de cartões de crédito"] & n8["Taxa de juros e influência em cada despesa"]
    n9["Domínios"] --> A & n1

    n9@{ shape: hex}
```

## Arquitetura

### Despesas

```mermaid
---
config:
  theme: redux
  layout: dagre
  look: neo
---
flowchart LR
    n14["Client WEB"] --> n1["Lambda - Expenses Service - Publisher"]
    n21["Client Mobile"] --> n1
    n1 --> n3@{ label: "Expenses<span style=\"color:\"> Topic</span>" }
    n3 --> n6["SQS - Transaction Queue"]
    n6 --> n22["Lambda - Expenses Service - Consumer"]
    n22 --> n18["Postgres - Expenses DB"] & n24(["External API - Interest Rate"])
    n18 --> n23["Lambda - DB Trigger - Alert Service"]

    n14@{ shape: rect}
    n1@{ shape: rect}
    n21@{ shape: rect}
    n3@{ shape: hex}
    n6@{ shape: h-cyl}
    n22@{ shape: rect}
    n18@{ shape: cyl}
    n23@{ shape: rect}
     n14:::Rose
     n21:::Sky
     n6:::Peach
     n18:::Sky
     n24:::Ash
    classDef Rose stroke-width:1px, stroke-dasharray:none, stroke:#FF5978, fill:#FFDFE5, color:#8E2236
    classDef Peach stroke-width:1px, stroke-dasharray:none, stroke:#FBB35A, fill:#FFEFDB, color:#8F632D
    classDef Sky stroke-width:1px, stroke-dasharray:none, stroke:#374D7C, fill:#E2EBFF, color:#374D7C
    classDef Ash stroke-width:1px, stroke-dasharray:none, stroke:#999999, fill:#EEEEEE, color:#000000
```

### Banking

```mermaid
---
config:
  theme: redux
  layout: dagre
  look: neo
---
flowchart LR
    n1["Lambda - Payments Service - Publiser"] --> n3["Payment Topic"]
    n3 --> n4["SQS - Antifraud Queue"] & n6["SQS - Transaction Queue"] & n15["SQS - Notification Queue"]
    n4 --> n5["Lambda - Antifraud Service - Consumer"]
    n6 --> n7["Lambda - Transaction Service - Consumer"]
    n14["Client WEB"] --> n1
    n15 --> n16["Lambda - Notification Service"]
    n5 --> n18["Postgres - Antifraud DB"] & n22["External - Bacen"]
    n5 -- status_update --> n3
    n7 --> n19["Postgres - Transaction DB"] & n22
    n7 -- status_update --> n3
    n16 --> n20["Postgres - Notification DB"] & n23["External - (SMS) AWS SNS"] & n24["External - (E-MAIL) SES or SEND GRID"]
    n16 -- status_update --> n3
    n21["Client Mobile"] --> n1

    n1@{ shape: rect}
    n3@{ shape: hex}
    n4@{ shape: h-cyl}
    n6@{ shape: h-cyl}
    n15@{ shape: h-cyl}
    n5@{ shape: rect}
    n7@{ shape: rect}
    n14@{ shape: rect}
    n16@{ shape: rect}
    n18@{ shape: cyl}
    n22@{ shape: rounded}
    n19@{ shape: cyl}
    n20@{ shape: cyl}
    n23@{ shape: rounded}
    n24@{ shape: rounded}
    n21@{ shape: rect}
     n4:::Peach
     n6:::Peach
     n15:::Peach
     n14:::Rose
     n18:::Sky
     n22:::Ash
     n19:::Sky
     n20:::Sky
     n23:::Ash
     n24:::Ash
     n21:::Sky
    classDef Rose stroke-width:1px, stroke-dasharray:none, stroke:#FF5978, fill:#FFDFE5, color:#8E2236
    classDef Peach stroke-width:1px, stroke-dasharray:none, stroke:#FBB35A, fill:#FFEFDB, color:#8F632D
    classDef Sky stroke-width:1px, stroke-dasharray:none, stroke:#374D7C, fill:#E2EBFF, color:#374D7C
    classDef Ash stroke-width:1px, stroke-dasharray:none, stroke:#999999, fill:#EEEEEE, color:#000000
```

## Ambientes Docker

### Ver mensagens no tópico

> http://localhost:8080

### Iniciar containers

```zsh
docker compose up
```

## Utils

- [Reset de tópico no Kafka](./docs/kafka/como_deletar_topico.md)
- [Teste de envio de pagamento](./docs/http/POST.http)
- [Doc do TCC](https://docs.google.com/document/d/1JjoOGEuK71dnkTWEYsoSsqkzM1umRoWZ/edit?usp=sharing&ouid=113161032001964584333&rtpof=true&sd=true)
