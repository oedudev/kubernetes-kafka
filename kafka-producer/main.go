package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func createTopic(topicName string, brokerAddress string) {
	// Criar um novo escritor Kafka
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokerAddress, topicName, 0)
	if err != nil {
		log.Fatalf("Erro ao conectar ao líder Kafka: %v", err)
	}
	defer conn.Close()
	// Listar tópicos existentes
	topics, err := conn.ReadPartitions()
	if err != nil {
		log.Fatalf("Erro ao ler partições: %v", err)
	}

	// Verificar se o tópico já existe
	topicExists := false
	for _, topic := range topics {
		if topic.Topic == topicName {
			topicExists = true
			break
		}
	}

	// Criar o tópico se ele não existir
	if !topicExists {
		log.Printf("Tópico %s não existe. Criando tópico...", topicName)
		topicConfig := kafka.TopicConfig{
			Topic: topicName,
		}
		err = conn.CreateTopics(topicConfig)
		if err != nil {
			log.Fatalf("Erro ao criar tópico: %v", err)
		}
		log.Printf("Tópico %s criado com sucesso", topicName)
	} else {
		log.Printf("Tópico %s já existe", topicName)
	}
}

func main() {

	// Definir o endereço do broker Kafka
	brokerAddress := "kafka:9092"

	// Definir o nome do tópico
	topicName := "my-topic"

	createTopic(topicName, brokerAddress)

	// Configuração do writer Kafka
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress}, // Endereço do broker Kafka
		Topic:    topicName,               // Nome do tópico
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	// Contexto com timeout
	ctx := context.Background()

	// Produzir mensagens a cada 1 segundo
	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {

		unixNow := time.Now().Unix()
		err := writer.WriteMessages(ctx,
			kafka.Message{
				Key:   []byte(t.String()),
				Value: []byte(fmt.Sprintf("%d", unixNow)),
			},
		)

		if err != nil {
			log.Printf("Erro ao escrever mensagem: %v\n", err)
			time.Sleep(time.Second * 20)
			continue
		}

		log.Printf("Mensagem enviada: %d\n", unixNow)
	}

}
