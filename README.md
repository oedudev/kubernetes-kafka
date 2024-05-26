
# Kafka Example Project

Este projeto é um exemplo de configuração de um ambiente Kafka no Kubernetes, com serviços Go para produção e consumo de mensagens. O projeto inclui um serviço `kafka-producer` que envia mensagens para um tópico Kafka e um serviço `kafka-consumer` que lê mensagens de todas as partições do tópico simultaneamente.

## Estrutura do Projeto

kafka-example/
├── kafka-consumer/
│ ├── main.go
│ ├── Dockerfile
│ └── k8s/
│ └── deployment.yaml
├── kafka-producer/
│ ├── main.go
│ ├── Dockerfile
│ └── k8s/
│ └── deployment.yaml
├── producer/
│ ├── main.go
│ ├── Dockerfile
│ └── k8s/
│ └── deployment.yaml
├── consumer/
│ ├── main.go
│ ├── Dockerfile
│ └── k8s/
│ └── deployment.yaml
├── nats/
│ └── k8s/
│ └── statefulset.yaml
├── kafka/
│ └── k8s/
│ ├── kafka-deployment.yaml
│ └── zookeeper-deployment.yaml
├── helm/
│ └── kafka-example/
│ ├── Chart.yaml
│ ├── values.yaml
│ ├── templates/
│ ├── producer-deployment.yaml
│ ├── consumer-deployment.yaml
│ ├── nats-deployment.yaml
├── skaffold.yaml
└── setup.sh

## Pré-requisitos

- [Docker](https://www.docker.com/)
- [Kubernetes](https://kubernetes.io/)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Helm](https://helm.sh/)
- [Skaffold](https://skaffold.dev/)

## Configuração do Ambiente

1. **Clone o Repositório:**
    git clone https://github.com/seu-usuario/kafka-example.git
    cd kafka-example

2. Torne o Script Executável:
    chmod +x setup.sh

3. Execute o Script de Configuração:

O script instalará os charts Helm necessários e iniciará o skaffold
    ./setup.sh

## Serviços

### Kafka Producer

O serviço `kafka-producer` envia mensagens para um tópico Kafka a cada segundo.

- **Código Fonte:** `kafka-producer/main.go`
- **Dockerfile:** `kafka-producer/Dockerfile`
- **Deployment Kubernetes:** `kafka-producer/k8s/deployment.yaml`

### Kafka Consumer

O serviço `kafka-consumer` lê mensagens de todas as partições de um tópico Kafka simultaneamente.

- **Código Fonte:** `kafka-consumer/main.go`
- **Dockerfile:** `kafka-consumer/Dockerfile`
- **Deployment Kubernetes:** `kafka-consumer/k8s/deployment.yaml`
