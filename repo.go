package main

import (
	"fmt"

	"gopkg.in/redis.v3"
)

type repo interface {
	GetPuzzleCount() int
}

type RedisRepo struct {
	client *redis.Client
}

var puzzlesKey string = "puzzles"
var listKey string = "puzzleList"

func NewRedisRepo(addr string) RedisRepo {
	fmt.Println("Connecting to redis on", addr)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return RedisRepo{client: client}
}

func (r RedisRepo) GetPuzzleCount() int {
	return int(r.client.SCard(listKey).Val())
}
