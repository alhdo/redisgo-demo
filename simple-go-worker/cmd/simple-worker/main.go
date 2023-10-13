package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/alhdo/simple-go-worker/core"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alhdo/simple-go-worker/api"
	"github.com/alhdo/simple-go-worker/config"
	"github.com/alhdo/simple-go-worker/cors"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Version Placeholder will be used at compile time
var Version = "0.0.0-src"

// Configuration Redis
var redisPool *redis.Pool

// Topic
const (
	JOB_TOPIC      = "jobTopic"
	RESPONSE_TOPIC = "jobResponseTopic"
)

func main() {

	v := false
	flag.BoolVar(&v, "v", false, "version")
	configFile := flag.String("c", "simple-go-worker.conf", "config file")
	flag.Parse()
	if v {
		fmt.Println("simple-worker", Version)
		os.Exit(0)
	}
	config := config.New(*configFile)
	redisConfig := core.RedisConfig{Addr: config.Redis.Host + ":" + config.Redis.Port}
	redisPool = core.InitRedisPool(redisConfig)

	pubSubservice := core.NewPubSubService(redisPool, RESPONSE_TOPIC, JOB_TOPIC)

	// Create a channel to receive message
	messages := make(chan string)

	// Start the subscriber in a go routine
	go pubSubservice.Subscribe(messages)

	// Process incoming message
	go func() {
		for message := range messages {
			response, _ := processJob(message, redisPool)
			pubSubservice.PublishMessage(response)
		}
	}()

	// Create request handler
	corsHandler := &cors.Handler{}
	fetchHandler := &api.FetchHandler{}

	r := mux.NewRouter()
	r.PathPrefix("/").Methods("OPTIONS").Handler(corsHandler)
	r.Path("/vlille").Methods("GET").Handler(fetchHandler)
	//go processJobs()
	log.Println("Worker started on ", config.Server.Address)
	log.Fatal(http.ListenAndServe(config.Server.Address, cors.AddCorsHeaders(r)))

}

func processJob(message string, pool *redis.Pool) (string, error) {
	// Fetch data
	stations := api.FetchVelibData()
	jsonData, err := json.Marshal(stations)
	if err != nil {
		fmt.Println("Error marshaling JSON data:", err)
		return "", err
	}
	rediskey := uuid.New().String()
	ttl := 60

	_, err = redisPool.Get().Do("SET", rediskey, string(jsonData), "EX", ttl)
	if err != nil {
		fmt.Println("Error setting value in Redis:", err)
		return "", err
	}
	jobData := core.JobResponse{
		Key:    rediskey,
		Status: core.Completed,
	}
	jobDataJSON, _ := json.Marshal(jobData)
	return string(jobDataJSON), nil
}
func initRedisPool(redisHost string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}
