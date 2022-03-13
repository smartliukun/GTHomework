package main

import (
	userservice "example/gthomework/work7/api"
	"example/gthomework/work7/user_server/internal/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// 1. new一个grpc的server
	rpcServer := grpc.NewServer()

	// 2. 将刚刚我们新建的ProdService注册进去
	userservice.RegisterUserServiceServer(rpcServer, services.UserServiceImpl)

	// 3. 新建一个listener，以tcp方式监听8082端口
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}

	//go initConsumer()
	//go initProducer()

	// 4. 运行rpcServer，传入listener
	_ = rpcServer.Serve(listener)
}

//func initConsumer() {
//
//	c, err := kafka.NewConsumer(&kafka.ConfigMap{
//		"bootstrap.servers": "localhost",
//		"group.id":          "myGroup",
//		"auto.offset.reset": "earliest",
//	})
//
//	if err != nil {
//		panic(err)
//	}
//
//	c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)
//
//	for {
//		msg, err := c.ReadMessage(-1)
//		if err == nil {
//			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
//		} else {
//			// The client will automatically try to recover from all errors.
//			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
//		}
//	}
//
//	c.Close()
//}
//
//func initProducer() {
//
//	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
//	if err != nil {
//		panic(err)
//	}
//
//	defer p.Close()
//
//	// Delivery report handler for produced messages
//	go func() {
//		for e := range p.Events() {
//			switch ev := e.(type) {
//			case *kafka.Message:
//				if ev.TopicPartition.Error != nil {
//					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
//				} else {
//					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
//				}
//			}
//		}
//	}()
//
//	// Produce messages to topic (asynchronously)
//	topic := "myTopic"
//	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
//		p.Produce(&kafka.Message{
//			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
//			Value:          []byte(word),
//		}, nil)
//	}
//
//	// Wait for message deliveries before shutting down
//	p.Flush(15 * 1000)
//}
