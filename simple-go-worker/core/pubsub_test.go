package core_test

import (
	"github.com/alhdo/simple-go-worker/core"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestMewPubSubService(t *testing.T) {
	// Set up a Redis pool for testing (you may want to use a mock Redis server for testing)
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	subChannel := "testSubChannel"
	pubChannel := "testSubChannel"

	// Create a PubSubService instance for testing
	ps := core.NewPubSubService(redisPool, pubChannel, subChannel)

	// Create a channel to receive messages
	messages := make(chan string, 1)

	// Start a goroutine to subscribe to the channel
	go func() {
		ps.Subscribe(messages)
	}()

	// Publish a test message
	testMessage := "Test Message"
	err := ps.PublishMessage(testMessage)
	if err != nil {
		t.Fatalf("Error publishing message: %v", err)
	}

	// Receive and validate the message
	receivedMessage := <-messages
	if receivedMessage != testMessage {
		t.Fatalf("Expected message '%s', but received '%s'", testMessage, receivedMessage)
	}
}
