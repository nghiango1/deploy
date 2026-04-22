package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nghiango1/deploy/webhook-handler/pkg/logger"
	"github.com/nghiango1/deploy/webhook-handler/pkg/webhook"
)

// Default http://127.0.0.1:3000/webhook
func main() {
	logger.SetupLogger()
	http.HandleFunc("/webhook", webhook.Handler)

	port := "3000"
	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
