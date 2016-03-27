package redis

import (
	"fmt"

	"github.com/mediocregopher/radix.v2/pool"
)

type RedisPool struct {
	*pool.Pool
}

func NewRedisPool(addr string, poolSize int) (*RedisPool, error) {
	pool, err := pool.New("tcp", addr, poolSize)
	if err != nil {
		fmt.Println("Unable to create pool")
		fmt.Println("Error: ", err.Error())
		return nil, err
	}

	return &RedisPool{pool}, nil
}
