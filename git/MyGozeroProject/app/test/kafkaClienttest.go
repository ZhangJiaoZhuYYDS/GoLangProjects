// @Author zhangjiaozhu 2023/10/3 23:59:00
package main

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// 设置 Kafka 连接配置
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// 创建 Kafka 消费者
	consumer, err := sarama.NewConsumer([]string{"192.168.31.5:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalf("Error closing Kafka consumer: %v", err)
		}
	}()

	// 订阅主题
	topic := "topic-beyond-like"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalf("Error closing partition consumer: %v", err)
		}
	}()

	// 处理收到的消息
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Received message: %s\n", string(msg.Value))
		case err := <-partitionConsumer.Errors():
			log.Printf("Error: %v\n", err.Err)
		case <-signals:
			return
		}
	}

}
