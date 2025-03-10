# 3 - Load Test

**Título da Decisão: Testes de estresse**

## 1. Contexto
Como o objetivo é testar alguns limites de arquiteturas asincrônas concorrentes, é muito importante
ter uma ferramenta para testes de estresse 

## 2. Decisão
Tomei a decição de colocar o K6 como ferramenta para testes de estresse.

Nunca fiz o uso da ferramenta, muito menos consegui ler a documentação para ter um bom contexto 
de quais os principais recursos.

Dito isso, irei descobrindo as utilidades e possibilidades da ferramenta de acordo com o tempo e uso
nesse projeto.

## **3. Consequências**  

**Positivas:**  
- Introdução de uma ferramenta de mercado amplamente utilizada para testes de carga, permitindo avaliações mais realistas da resiliência do sistema.  
- Possibilidade de automatizar testes de estresse, ajudando na detecção antecipada de gargalos e problemas de performance.  
- Melhor compreensão dos limites da arquitetura assíncrona concorrente, possibilitando ajustes e otimizações conforme necessário.  
- Facilidade de integração com outras ferramentas de monitoramento e observabilidade, como Prometheus e Grafana.  

**Negativas:**  
- Curva de aprendizado inicial, já que a ferramenta ainda não foi estudada profundamente.  
- Possível necessidade de ajustes no ambiente de testes para suportar a execução dos testes de carga de forma eficiente.  
- Risco de interpretação errônea dos resultados devido à falta de experiência inicial com a ferramenta.  
- Investimento de tempo adicional para aprendizado e experimentação antes de obter benefícios concretos.