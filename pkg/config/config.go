package config

var config Configuration

type Configuration struct {
	MongoURI string
}

func MongoURI() string {
	return config.MongoURI
}
