# Concurrency Simulator

## Descrição 

Api com algum recurso simulado para simular um cenário de concorrência.

## Requisitos

### Funcionais

| ID       | Descrição                                                                 |
|----------|---------------------------------------------------------------------------|
| RF-001   | deve receber um post e criar um regístro de pagamento                     |

### Não funcionais

| ID       | Descrição                                                                 |
|----------|---------------------------------------------------------------------------|
| RNF-001  | processar em menos de 100 ms                                              |
| RNF-002  | precisa ser feito em Golang                                               |
| RNF-003  | simular um ambiente de microsserviços                                     |
| RNF-004  | usar postgres como banco de dados                                         |
| RNF-005  | usar kafka como sistema de filas                                          |
| RNF-006  | ter um sistema de tracking e observabilidade                              |
| RNF-007  | utilizar event sourcing                                                   |

## ADR

**Título da Decisão: Sistema simulador de concorrência**  

### 1. Contexto
Estou tentando criar um ambiente simulado de concorrência, visto que isso é um cenário muito comum em aplicações com milhares de requisições.

### 2. Decisão
Como o objetivo desse projeto é apenas testar minhas capacidades técnicas em elaborar soluções escaláveis, as escolhas de técnologias foram totalmente arbitrárias.

### 3. Consequências
**Positivas:**
- O simulador de concorrência permitirá uma análise prática da escalabilidade e do comportamento do sistema sob carga, possibilitando entender melhor as limitações e o desempenho da API em cenários de alta concorrência.
- A implementação em Golang, com PostgreSQL e Kafka, simula um ambiente de microsserviços e traz benefícios como:
  - **Baixa latência:** Golang é conhecido por sua eficiência em lidar com concorrência e baixa latência, o que ajuda a atingir o requisito de processamento em menos de 100 ms.
  - **Escalabilidade:** O uso de Kafka como sistema de filas facilita a distribuição e o balanceamento de carga entre serviços, permitindo que o sistema escale horizontalmente.
  - **Consistência de dados:** O PostgreSQL garante a durabilidade e consistência dos registros de pagamento, fundamentais para testes em cenários que exigem integridade transacional.

**Negativas:**
- **Complexidade de setup:** A necessidade de configurar Kafka, PostgreSQL e microsserviços, mesmo em um ambiente simulado, pode resultar em maior tempo e esforço para a inicialização do projeto.
- **Custo de manutenção:** Manter e monitorar uma arquitetura simulada de microsserviços, mesmo em um ambiente de testes, pode gerar complexidade adicional, especialmente com a adição de tracking e observabilidade.
- **Limitações de simulação:** Apesar do uso de tecnologias robustas, o simulador pode não capturar todas as nuances de um ambiente de produção real, especialmente quando se trata de comportamento de usuários e carga imprevisível.

**Riscos:**
- O simulador pode se tornar mais complexo do que o esperado, desviando o foco dos objetivos iniciais e exigindo mais esforço para atender os requisitos de baixa latência e alta disponibilidade.
- A escolha arbitrária de tecnologias pode não refletir o ambiente de produção que será usado futuramente, limitando a transferência dos aprendizados para um sistema real.
