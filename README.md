# Desafio Técnico DevOps / Infraestrutura - Korp

Este repositório contém a solução para o desafio técnico da Korp, englobando a containerização de um microsserviço em Go, configuração de proxy reverso, observabilidade/monitoramento e automação do deploy utilizando Infraestrutura como Código (IaC).

---

## 🏗️ Arquitetura do Projeto

A infraestrutura foi desenhada para ser totalmente resiliente, isolada e automatizada, dividida em três pilares principais:

1. **Aplicações & Proxy (Porta 80):** - Um servidor HTTP em **Go** utilizando Dockerfile otimizado via *Multi-stage Build*.
   - Um servidor **NGINX** atuando como Proxy Reverso, recebendo as requisições externas na porta `80` e direcionando de forma segura para o container Go através de uma rede isolada (`korp-network`).

2. **Monitoramento & Observabilidade:**
   - **Prometheus (Porta 9090):** Configurado para realizar o *scrape* (coleta) periódico das métricas expostas pela aplicação Go na rota `/metrics`.
   - **Grafana (Porta 3000):** Integrado ao Prometheus como Fonte de Dados (*Data Source*) para a criação de dashboards e visualização analítica do comportamento do servidor.

3. **Automação (IaC):**
   - Um Playbook do **Ansible** encarregado de validar o ambiente local e orquestrar a execução do ecossistema de containers sem necessidade de intervenção manual passo a passo.

---

## 📁 Estrutura de Arquivos

```text
projeto-korp/
├── main.go               # Código fonte do servidor HTTP em Go
├── go.mod                # Módulo de dependências do Go
├── go.sum                # Checksums das dependências do Go
├── Dockerfile            # Build otimizado em estágios (Multi-stage)
├── docker-compose.yml    # Orquestração dos serviços da infraestrutura
├── nginx.conf            # Configuração do Proxy Reverso NGINX
├── prometheus.yml        # Configuração de coleta do Prometheus
└── playbook.yml          # Automação do Deploy com Ansible
