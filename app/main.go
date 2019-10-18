package main

import (
	"fmt"
	"net/http"

	"github.com/omaressameldin/feedback-ninja/app/pkg/confirmation"
	"github.com/omaressameldin/feedback-ninja/app/pkg/env"
	"github.com/omaressameldin/feedback-ninja/app/pkg/feedback"
)

func main() {
	env.ValidateEnvKeys()

	port := env.GetPort()
	http.HandleFunc("/", feedback.Handler)
	http.HandleFunc("/confirmation", confirmation.Handler)

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
