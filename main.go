package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	nsqAddress = "nsqd:4150"
	topic      = "example_topic"
	channel    = "example_channel"
)

func main() {
	// Setup NSQ producer
	producer, err := nsq.NewProducer(nsqAddress, nsq.NewConfig())
	if err != nil {
		log.Fatal("Error creating NSQ producer:", err)
	}
	defer producer.Stop()

	go func() {
		for {
			err := producer.Publish(topic, []byte("Hello World!"))
			if err != nil {
				log.Fatal("Error publishing to NSQ:", err)
			}
			fmt.Println("Published message to NSQ")
			time.Sleep(1 * time.Second)
		}
	}()

	// Setup NSQ consumer
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal("Error creating NSQ consumer:", err)
	}

	// Define NSQ message handler
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Printf("Received message: %s\n", message.Body)
		return nil
	}))

	// Connect consumer to NSQ
	err = consumer.ConnectToNSQD(nsqAddress)
	if err != nil {
		log.Fatal("Error connecting consumer to NSQ:", err)
	}

	//Wait for interrupt signal to gracefully stop the consumer
	var wg sync.WaitGroup
	wg.Add(1)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-signalCh:
			consumer.Stop()
			wg.Done()
		}
	}()

	wg.Wait()
}
