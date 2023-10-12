package main

import (
	"encoding/json"
	"flag"
	"fmt"
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
	fmt.Println(config.Redis.Host + ":" + config.Redis.Port)
	redisPool = initRedisPool(config.Redis.Host + ":" + config.Redis.Port)
	// Create request handler
	corsHandler := &cors.Handler{}
	fetchHandler := &api.FetchHandler{}

	r := mux.NewRouter()
	r.PathPrefix("/").Methods("OPTIONS").Handler(corsHandler)
	r.Path("/vlille").Methods("GET").Handler(fetchHandler)
	go processJobs()
	log.Println("Worker started on ", config.Server.Address)
	log.Fatal(http.ListenAndServe(config.Server.Address, addCorsHeaders(r)))

}

type JobData struct {
	TaskID string `json:"taskID"`
	Data   string `json:"data"`
}

func processJobs() {
	conn := redisPool.Get()
	defer conn.Close()

	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe("jobTopic")

	pubConn := redisPool.Get()
	defer pubConn.Close()

	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			// Nouveau job reçu
			fmt.Printf("Nouveau job reçu : %s\n", v.Data)

			// Fetch station
			//fh := api.FetchHandler{}
			stations := api.FetchVelibData()
			jsonData, err := json.Marshal(stations)
			if err != nil {
				fmt.Println("Error marshaling JSON data:", err)
				return
			}
			rediskey := uuid.New().String()
			ttl := 60

			_, err = redisPool.Get().Do("SET", rediskey, string(jsonData), "EX", ttl)
			if err != nil {
				fmt.Println("Error setting value in Redis:", err)
				return
			}

			// Stockez les données traitées dans Redis
			// _, err := conn.Do("SET", taskID, "Données traitées", "EX", 60)
			// if err != nil {
			//     log.Println("Échec de la mise en cache des données dans Redis")
			// }
			// ...

			// Envoyez un message de confirmation
			jobData := JobData{
				TaskID: rediskey,
				Data:   "Données traitées",
			}
			jobDataJSON, _ := json.Marshal(jobData)
			_, err = pubConn.Do("PUBLISH", "jobResponseTopic", jobDataJSON)
			if err != nil {
				fmt.Println("Job publish error ", err.Error())
			}

		case redis.Subscription:
			fmt.Printf("Abonnement à un canal : %s %s %d\n", v.Kind, v.Channel, v.Count)
		case error:
			fmt.Printf("Erreur Redis Pub/Sub : %v\n", v)
			return
		}
	}
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
func addCorsHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow GET, POST, PUT, DELETE methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

		// Allow the "Content-Type" header
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Allow credentials (if needed)
		// w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests (OPTION method)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Continue to the next handler
		handler.ServeHTTP(w, r)
	})
}
