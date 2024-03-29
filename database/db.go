package database

import (
	"eviecoin/environment"
	"eviecoin/item_database"
	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
	"os"
)

type Database struct {
	Users users
	Items item_database.ItemDatabase
}

func NewDatabase() Database {
	color.Blue("[database] Getting a redis eviecoin ready...")
	getClient()

	return Database{
		Items: item_database.NewManager(),
		Users: users{
			GetUser:    getUser,
			SetUser:    setUser,
			AddBalance: addBalance,
			AddItem:    addItem,
		},
	}
}

var client *rejson.Handler

func getClient() *rejson.Handler {
	if client != nil {
		return client
	}

	color.Blue("[database] Creating a new redis instance...")

	client = rejson.NewReJSONHandler()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       environment.GetIntWithDefault("REDIS_DB", 0),
	})

	client.SetGoRedisClient(rdb)

	return client
}
