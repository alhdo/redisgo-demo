<h1 align="center">Real-time Velo Station Availability Monitoring</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-0.0.1-blue.svg?cacheSeconds=2592000" />
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="https://twitter.com/castroalhdo" target="_blank">
    <img alt="Twitter: castroalhdo" src="https://img.shields.io/twitter/follow/castroalhdo.svg?style=social" />
  </a>
</p>

> Demonstration of using redis pub sub cannal with go an vue js

Welcome to the Real-time Velo Station Availability Monitoring repository! This project consists of three main components that work together to provide real-time updates on the availability of bikes at Velo stations.


## Components

### 1. Go Worker (Worker Component)
The Go Worker is a program that acts as a subscriber to a Redis Pub/Sub channel. Its main purpose is to fetch a list of available places for a Velo station. This component plays a crucial role in collecting data and preparing it for further processing by the Node.js API.


### 2. Node.js API (API Component)
The Node.js API is responsible for subscribing to the Redis Pub/Sub channel, receiving signal of data processed by the Go Worker, and retrieve this data in a Redis database. It also takes care of publishing the real-time updates to a WebSocket.

### 3. Vue.js Front-End Application (Front-End Component)

The Vue.js Front-End Application is the user interface that visualizes the real-time data about available bikes at Velo stations. It subscribes to the WebSocket to display this information to the end users in an intuitive and user-friendly manner.

## Getting Started

### Run With Docker
The components are orchestrated with docker. To Run all the 3 components make sure you have docker and compose install and run :

```sh
cd redisgo-demo
docker-compose up -d
```

The front-end will be running on port 8000

[http://localhost:8000](http://localhost:8000)

You can run this project by running each of the component separately also. Below is the step and configuration needed to run them independently:


1. Clone the repository to your local machine.
2. Set up the necessary dependencies for each component. You may need to install Go, Node.js, and Vue.js.
3. Configure the connection to your Redis server in the Go Worker and Node.js API.
4. Start the Node.js API to process the data and publish it to the WebSocket.
5. Run the Go Worker to start collecting data.
6. Launch the Vue.js Front-End Application to see real-time bike availability updates.

#### API Config
```
API_PORT=3000
REDIS_HOST="localhost"
REDIS_PORT=6379
```

#### GO Worker Config

```
[server]
## Defines the IP address and port the HTTP server listens on.
## Defaults: "0.0.0.0:8080"
address = "0.0.0.0:8080"

[redis]
## Defines the IP address and port of redis instance
host = "127.0.0.1"
port = "6379"
```

#### Vuejs Config

```
VITE_API_URL="http://localhost:3000"
```

## Configuration


The project consist of 3 components :
  - `simple-go-worker` A GO worker
  - `api` A NodeJS API
  - `vlille` A front-end using Vuejs

## Author

👤 **Alhdo**

* Twitter: [@castroalhdo](https://twitter.com/castroalhdo)
* Github: [@alhdo](https://github.com/alhdo)

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_