package env

import (
	"log"

	"github.com/omaressameldin/feedback-ninja/app/internal/utils"
)

func ValidateEnvKeys() {
	GetToken()
	GetPort()
	GetFeedbackChannelID()
}

func GetToken() string {
	token, err := utils.GetEnv(tokenKey)
	if err != nil {
		log.Fatal(err)
	}

	return token
}

func GetPort() string {
	port, err := utils.GetEnv(portKey)
	if err != nil {
		log.Fatal(err)
	}

	return port
}

func GetFeedbackChannelID() string {
	channelID, err := utils.GetEnv(feedbackChannelKey)
	if err != nil {
		log.Fatal(err)
	}

	return channelID
}
