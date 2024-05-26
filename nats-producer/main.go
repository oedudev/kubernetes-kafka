package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	log.Print("Iniciando serviço")

	var nc *nats.Conn
	var err error
	for {
		// Endereços dos servidores NATS no cluster
		natsURLs := "nats://nats-0.nats:4222,nats://nats-1.nats:4222,nats://nats-2.nats:4222,nats://nats-3.nats:4222,nats://nats-4.nats:4222"

		// Configurar opções de conexão
		opts := []nats.Option{
			nats.Name("producer"),
			nats.Timeout(10 * time.Second),
			nats.ReconnectWait(5 * time.Second),
			nats.MaxReconnects(-1), // Tentar reconectar indefinidamente
		}

		// Conectar ao cluster NATS
		nc, err = nats.Connect(natsURLs, opts...)
		if err != nil {
			log.Printf("Erro ao conectar ao cluster NATS: %v\n", err)
			continue
		}
		defer nc.Close()
		break
	}

	log.Println("Conectado ao cluster NATS")
	for {
		unixDate := time.Now().Unix()
		// Publicar uma mensagem
		msg := []byte(fmt.Sprintf("%d", unixDate))

		if err := nc.Publish("test", msg); err != nil {
			fmt.Printf("error: %s\n", err.Error())
		}
		log.Println("mensagem inserida na fila NATS")
		time.Sleep(time.Second * 30)
	}

}
