package config_test

import (
	"github.com/alhdo/simple-go-worker/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	config := config.New("../simple-go-worker.conf.example")
	assert.Equal(t, "127.0.0.1:8080", config.Server.Address, "error getting [server] address")
	assert.Equal(t, "6379", config.Redis.Host, "error getting [redis] port")
	assert.Equal(t, "127.0.0.1", config.Redis.Host, "error getting [redis] host")
}
