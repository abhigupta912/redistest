package producer

import "github.com/abhigupta912/redistest/test01/redis"

type QProducer struct {
	pool      *redis.RedisPool
	queueName string
}

func NewQProducer(pool *redis.RedisPool, qName string) QProducer {
	return QProducer{pool, qName}
}

func (producer QProducer) ProduceMsg(content string) error {
	response := producer.pool.Cmd("LPUSH", producer.queueName, content)
	if response.Err != nil {
		return response.Err
	}
	return nil
}
