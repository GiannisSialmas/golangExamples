package rediscache

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

//	ENV Variables to configure the redis connection
//	REDIS_HOST=null 	- Redis url to connect to. Program terminates if not given
//	REDIS_PORT=6379 	- Redis instance port
//  REDIS_PASSWORD=""	- Password to use for connection to the redis instance

var RedisClient *redis.Client
var CacheActive = false

func init() {

	// Initialize the cache as false because we dont know if we have it yet
	redisConnectionString := getRedisAdress()
	redisPassword := os.Getenv("REDIS_PASSWORD")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConnectionString,
		Password: redisPassword,
		DB:       0, // use default DB
	})

	log.Println("Trying to connect to single redis instance at ", redisConnectionString)
	_, err := RedisClient.Ping().Result()

	if err != nil {
		log.Println("Error while trying to ping redis instance:", err)
	} else {
		log.Println("Succesfully ping the redis instance")
		CacheActive = true
	}

	//begin goroutine to periodically check for the cache health
	go func() {
		for range time.Tick(time.Second) {
			checkForConnectivity()
		}
	}()

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
		log.Printf("No port given for redis instance. Defaulting to port %s\n", port)
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalln("Invalid port", port)
	}

	return fmt.Sprintf("%s:%s", host, port)
}

func checkForConnectivity() {
	_, err := RedisClient.Ping().Result()

	if err != nil {
		log.Println("Error while trying to ping redis instance:", err)
		CacheActive = false
	} else {
		log.Println("Succesfully ping the redis instance")
		CacheActive = true
	}
}

// // Simple example that is not being used anywhere
// func simpleSetGetExample(redisClient *redis.Client) {
// 	err := redisClient.Set("key", "value", 0).Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	val, err := redisClient.Get("key").Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("key", val)

// 	val2, err := redisClient.Get("key2").Result()
// 	if err == redis.Nil {
// 		fmt.Println("key2 does not exist")
// 	} else if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("key2", val2)
// 	}
// }

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
