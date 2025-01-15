package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	mongoCS := os.Getenv("MONGO_STRING")

	config.MongoString = mongoCS

	return nil

}
