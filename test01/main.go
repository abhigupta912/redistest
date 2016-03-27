package main

import (
	"bytes"
	"fmt"
	"strconv"
	"sync"

	"github.com/abhigupta912/redistest/test01/consumer"
	"github.com/abhigupta912/redistest/test01/producer"
	"github.com/abhigupta912/redistest/test01/redis"
)

func main() {
	pool, err := redis.NewRedisPool("localhost:6379", 10)
	if err != nil {
		panic(err)
	}

	qName := "testq"

	// Initialize producers
	producerCount := 3

	var producerWg sync.WaitGroup
	producerWg.Add(producerCount)

	for pIdx := 1; pIdx <= producerCount; pIdx++ {
		go func(p *redis.RedisPool, q string, i int) {
			InitProducer(p, q, i)
			producerWg.Done()
		}(pool, qName, pIdx)
	}

	producerWg.Wait()

	// Initialize consumers
	consumerCount := 5

	var consumerWg sync.WaitGroup
	consumerWg.Add(consumerCount)

	for cIdx := 1; cIdx <= consumerCount; cIdx++ {
		go func(p *redis.RedisPool, q string, i int) {
			InitConsumer(p, q, i)
			consumerWg.Done()
		}(pool, qName, cIdx)
	}

	consumerWg.Wait()
}

func InitConsumer(pool *redis.RedisPool, qName string, index int) {
	c := consumer.NewQConsumer(pool, qName)
	msg, err := c.ConsumeMsg()
	if err == nil {
		fmt.Printf("Consumer %d recd message %s from queue %s\n", index, msg, qName)
	}
}

func InitProducer(pool *redis.RedisPool, qName string, index int) {
	c := producer.NewQProducer(pool, qName)

	message := bytes.NewBufferString("Message")
	message.WriteString(strconv.Itoa(index))
	msg := message.String()

	fmt.Printf("Producer %d posting message %s to queue %s\n", index, msg, qName)
	c.ProduceMsg(msg)
}
