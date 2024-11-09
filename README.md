# Concurrency Simulator

## Descrição 

Api com algum recurso simulado para simular um cenário de concorrência.

## Requisitos

### Funcionais

| ID       | Descrição                                                                 |
|----------|---------------------------------------------------------------------------|
| RF-001   | deve receber um post e adicionar um item ao carrinho de compras           |

### Não funcionais

| ID       | Descrição                                                                 |
|----------|---------------------------------------------------------------------------|
| RNF-001  | processar em menos de 100 ms                                              |
| RNF-002  | precisa ser feito em Golang                                               |
| RNF-003  | simular um ambiente de microserviços                                      |
| RNF-004  | usar postgres como banco de dados                                         |
| RNF-005  | usar kafka como sistema de filas                                          |
| RNF-006  | ter um sistema de tracking e observabilidade                              |
