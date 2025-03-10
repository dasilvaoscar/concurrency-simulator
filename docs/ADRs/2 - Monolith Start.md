# 2 - Monolith Start

**Título da Decisão: Início como monolito**

## 1. Contexto
Começar já levando em consideração as complexidades de assincrônicidade e comunicação entre microsserviços pode tirar a atenção da construção inicial do projeto

## 2. Decisão
Vou começar a aplicação como monolito na Core API, vendo que o objetivo final é a transformação disso em microsserviços, é **necessário** estar bem definido os limites de cada parte do sistema.
O desacoplamento entre a base do código é uma prioridade, e é necessário existir uma boa suit de testes para garantir uma migração suave.

## 3. Consequências
**Positivas:**
- Migração para microsserviços deve ser bem tranquila
- A base do código deve estar muito organizada
- Boa cobertura de testes

**Negativas:**
- O tempo para maturação e execução do software será maior
- A complexidade inicial para separação de domínios e desacoplamento do código a curto prazo é alta
