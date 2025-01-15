package config

var config Configuration

type Configuration struct {
	MongoString string
}

func MongoString() string {
	return config.MongoString
}
