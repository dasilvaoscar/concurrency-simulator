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
| RNF-008  | implementar setup docker para iniciar os serviços                         |

## ADR (Architectural Decision Record)
- [1 - Project Start](./docs/ADRs/1%20-%20Project%20Start.md)
- [2 - Monolith Start](./docs/ADRs/2%20-%20Monolith%20Start.md)
- [3 - Load Test](./docs/ADRs/3%20-%20Load%20Test.md)

## High Level Architecture

- [High Level Architecture Link](https://miro.com/welcomeonboard/Ymx1M214YVEyTHpNU3BFYmVHSXV0bEVNeDhvWU10allDUjJ1Smc4eGlOcjljbEZBRldETFJrbFd1WGRZUUtVMlhRTW54Ujd5UEtEQ3BsbVFxcGo4R1lmd0xrMTVwc0ljUkQ2OU9lU2x6T2Y3RUtZczJpZGQzTStuY0l2TGZ6L0chZQ==?share_link_id=599870259324)

![Captura de tela de 2025-03-09 22-33-51](https://github.com/user-attachments/assets/7de10d44-f03c-4f0a-93dd-3ebeb0c9d2ec)
![Captura de tela de 2024-11-10 22-12-16](https://github.com/user-attachments/assets/c9b959c5-e8ba-4f42-a86f-1a942a10b476)


## Ambientes Docker


- Para iniciar e criar as imagens, use o comando `docker compose up --build`
- Para iniciar todos os serviços, use o comando  `docker compose up`
- Para iniciar um serviço especifico, use o comando  `docker compose up [nome do serviço] `

| Service name      | Port |
|-------------------|------|
| antifraud-svc     | 8081 |
| core-svc          | 8082 |
| customer-svc      | 8083 |
| notification-svc  | 8084 |
| transaction-svc   | 8085 |


