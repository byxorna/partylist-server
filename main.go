package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/byxorna/partylist-server/web"
	"gopkg.in/redis.v3"
)

var (
	httpPort  int
	redisPort int
	redisHost string
)

func init() {
	flag.IntVar(&httpPort, "port", 8000, "HTTP port")
	flag.StringVar(&redisHost, "redis-host", "localhost", "Redis host")
	flag.IntVar(&redisPort, "redis-port", 6379, "Redis port")
	flag.Parse()
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(pong, err)
	}
	log.Info("Connection to redis verified ", client, pong)

	router := web.New(client)
	log.Infof("Starting webserver on %s", fmt.Sprintf(":%d", httpPort))
	if err := router.Run(fmt.Sprintf(":%d", httpPort)); err != nil {
		log.Fatal(err)
	}
}
