# Concurrency Simulator

## Descrição 

Api com algum recurso simulado para simular um cenário de concorrência.

## Requisitos

### Funcionais

| ID       | Descrição                                                                 |
|----------|---------------------------------------------------------------------------|
| RF-001   | deve receber um post e criar um regístro de pagamento                     |
| RF-002   | enviar push notifications com o status do pagamento                        |

### Não funcionais

| ID       | Descrição                                                                 |
|----------|---------------------------------------------------------------------------|
| RNF-001  | processar em menos de 100 ms                                              |
| RNF-002  | precisa ser feito em Golang                                               |
| RNF-003  | simular um ambiente de microsserviços                                     |
| RNF-004  | usar postgres como banco de dados SQL, e dynamo NoSQL                     |
| RNF-005  | usar kafka como sistema de filas                                          |
| RNF-006  | ter um sistema de tracking e observabilidade                              |
| RNF-007  | utilizar event sourcing                                                   |

## ADR
- [1 - Project Start](./docs/ADRs/1%20-%20Project%20Start.md)
- [2 - Monolith Start](./docs/ADRs/2%20-%20Monolith%20Start.md)

## High Level Architecture

- [High Level Architecture Link](https://github.com/user-attachments/assets/2dbdcb09-bdf2-4fbc-9947-41a259d432d1)

![Captura de tela de 2024-11-10 22-11-03](https://github.com/user-attachments/assets/9ad89638-0c65-44a8-ba72-fd0c898898a3)
![Captura de tela de 2024-11-10 22-12-16](https://github.com/user-attachments/assets/c9b959c5-e8ba-4f42-a86f-1a942a10b476)
