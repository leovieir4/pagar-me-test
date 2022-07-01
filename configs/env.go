package configs

import (
	"os"

	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

func GetNeo4JURI() string {
	err := godotenv.Load("./.env")
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			},
		).Debug("Not loaded .env file")
	}

	return os.Getenv("NEO4J_URI")

}

func GetNeo4jUserName() string {
	err := godotenv.Load("./.env")
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			},
		).Debug("Not loaded .env file")
	}

	return os.Getenv("NEO4J_USERNAME")

}

func GetNeo4JPassword() string {
	err := godotenv.Load()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			},
		).Debug("Not loaded .env file")
	}
	return os.Getenv("NEO4J_PASSWORD")
}
