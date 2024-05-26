package main

import (
	"context"
	"log"
	"sync"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Definir o endereço do broker Kafka
	brokerAddress := "kafka:9092"
	topicName := "my-topic"

	// Conectar ao Kafka para obter a lista de partições
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokerAddress, topicName, 0)
	if err != nil {
		log.Fatalf("Erro ao conectar ao líder Kafka: %v", err)
	}
	defer conn.Close()

	// Obter a lista de partições para o tópico
	partitions, err := conn.ReadPartitions(topicName)
	if err != nil {
		log.Fatalf("Erro ao ler partições: %v", err)
	}

	// Criar um WaitGroup para aguardar a leitura de todas as partições
	var wg sync.WaitGroup

	// Função para ler mensagens de uma partição
	readPartition := func(partition int) {
		defer wg.Done()

		// Configuração do leitor Kafka para a partição específica
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:   []string{brokerAddress},
			Topic:     topicName,
			Partition: partition,
		})
		defer reader.Close()

		log.Printf("Iniciando leitura da partição %d", partition)

		// Loop para consumir mensagens
		for {
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Printf("Erro ao ler mensagem da partição %d: %v", partition, err)
				return
			}
			log.Printf("Mensagem recebida da partição %d: %s", partition, string(msg.Value))
		}
	}

	// Iniciar uma goroutine para cada partição
	for _, p := range partitions {
		wg.Add(1)
		go readPartition(p.ID)
	}

	// Aguardar a leitura de todas as partições
	wg.Wait()
}
