package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	mongoURI := os.Getenv("MONGO_URI")

	config.MongoURI = mongoURI

	return nil

}
