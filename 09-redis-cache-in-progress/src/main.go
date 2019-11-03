package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
)

// var listenAdress string

func main() {

	redisConnectionString := getRedisAdress()
	fmt.Println("Redis instance is at", redisConnectionString)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisConnectionString,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)

	// simpleSetGetExample(redisClient)
	cacheExample(redisClient)
}

// Calculates the listen adress or the HTTP server from environmental variables
func getRedisAdress() string {

	// Get redis url from env variable
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		log.Fatalln("No redis host given port")
	}

	// Get redis port from env variable
	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = "6379"
		fmt.Printf("No port given for redis instance. Defaulting to port %s\n", port)
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalln("Invalid port", port)
	}

	return fmt.Sprintf("%s:%s", host, port)
}

func simpleSetGetExample(redisClient *redis.Client) {
	err := redisClient.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisClient.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := redisClient.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func cacheExample(redisClient *redis.Client) {
	codec := &cache.Codec{
		Redis: redisClient,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	type Object struct {
		FirstName string
		Age       int
	}

	key := "Sialmas"
	obj := &Object{
		FirstName: "giannis",
		Age:       42,
	}

	codec.Set(&cache.Item{
		Key:        key,
		Object:     obj,
		Expiration: time.Hour,
	})

	var wanted Object
	if err := codec.Get(key, &wanted); err == nil {
		fmt.Println("Retrieved:", wanted)
	}
}
