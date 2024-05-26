#!/bin/bash

# Adicionar repositórios Helm
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo add hashicorp https://helm.releases.hashicorp.com

# Atualizar repositórios Helm
helm repo update

# Instalar o Consul
helm install consul hashicorp/consul --set global.name=consul

# Instalar o Kafka com valores personalizados
cat <<EOF >kafka-values.yaml
replicaCount: 3
config:
  log.dirs: "/opt/bitnami/kafka/data"
  auto.create.topics.enable: "true"
  num.partitions: 3
EOF

helm install kafka bitnami/kafka -f kafka-values.yaml

# Remover o arquivo temporário kafka-values.yaml
rm kafka-values.yaml

# Rodar o Skaffold
skaffold dev --tail=false
