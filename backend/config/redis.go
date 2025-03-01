package config

import(
	"github.com/go-redis/redis"
	"log"
	"exchangeapp/global"
	
)

func initRedis(){
	RedisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB:0,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil{
		log.Fatalf("Error connecting to redis, %v", err)
	}
	global.RedisDB = RedisClient
}