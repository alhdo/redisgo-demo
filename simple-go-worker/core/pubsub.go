package core

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type PubSubService struct {
	pool       *redis.Pool
	subChannel string
	pubChannel string
}

func NewPubSubService(pool *redis.Pool, pubChannel string, subChannel string) *PubSubService {
	return &PubSubService{
		pool:       pool,
		pubChannel: pubChannel,
		subChannel: subChannel,
	}
}

// Subscribe to redis channel
func (ps *PubSubService) Subscribe(messages chan string) {
	conn := ps.pool.Get()
	defer conn.Close()

	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe(ps.subChannel)
	//conn.Flush()
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			// New job received
			fmt.Printf("Got a new job : %s\n", v.Data)
			messages <- string(v.Data)
		case redis.Subscription:
			fmt.Printf("Subscribtion to a channel : %s %s %d\n", v.Kind, v.Channel, v.Count)
		case error:
			fmt.Printf("Error Redis Pub/Sub : %v\n", v)
		}
	}
}

// Publish a message to channel

func (ps *PubSubService) PublishMessage(message string) error {
	conn := ps.pool.Get()
	defer conn.Close()
	_, err := conn.Do("PUBLISH", ps.pubChannel, message)
	if err != nil {
		fmt.Println("Message publish error ", err.Error())
		return err
	}
	return nil
}
