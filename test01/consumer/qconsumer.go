package consumer

import "github.com/abhigupta912/redistest/test01/redis"

type QConsumer struct {
	pool      *redis.RedisPool
	queueName string
}

func NewQConsumer(pool *redis.RedisPool, qName string) QConsumer {
	return QConsumer{pool, qName}
}

func (consumer QConsumer) ConsumeMsg() (string, error) {
	response := consumer.pool.Cmd("BRPOP", consumer.queueName, 5)
	if response.Err != nil {
		return "", response.Err
	}

	responseArr, err := response.Array()
	if err != nil {
		return "", err
	}

	message, err := responseArr[1].Str()
	if err != nil {
		return "", err
	}

	return message, nil
}
