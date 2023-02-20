package main

import (
	"Getir/internal/handlers"
	cache2 "Getir/internal/repositories/cache"
	storage2 "Getir/internal/repositories/storage"
	"Getir/internal/services"
	"log"
	"net/http"
)

func main() {
	config, err := GetConfig()
	if err != nil {
		log.Fatalf("loading config has failed. err: %s\n", err.Error())
	}

	redisClient, err := cache2.NewRedis(&cache2.RedisParams{
		Address:  config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
	if err != nil {
		log.Fatalf("create redis instance has failed. err: %s\n", err.Error())
	}
	cacheRepo := cache2.NewRepo(redisClient)

	mongoCollection, err := storage2.NewMongo(
		&storage2.MongoParams{
			Address:    config.MongoDB.Address,
			Database:   config.MongoDB.Database,
			Collection: config.MongoDB.Collection,
		})
	if err != nil {
		log.Fatalf("create mongo instance has failed. err: %s\n", err.Error())
	}
	storageRepo := storage2.NewRepo(mongoCollection)

	s := services.NewServices(storageRepo, cacheRepo)
	h := handlers.NewHandlers(s)

	mux := http.NewServeMux()
	mux.HandleFunc("/fetch", h.Fetch)
	mux.HandleFunc("/in-memory", h.InMemory)
	log.Fatal(http.ListenAndServe(config.Server.Address, mux))
}
