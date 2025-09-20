package main

import (
	"fmt"
	"log"
	"net/http"

	"go-gen-one/server"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("SBOM Test Project with Random Go Libraries")

	// UUID
	id := uuid.New()
	fmt.Println("UUID:", id)

	// Logrus
	logger := logrus.New()
	logger.Info("Logrus initialized")

	// Viper (config)
	viper.SetDefault("port", 8080)
	port := viper.GetInt("port")
	fmt.Println("Server port:", port)

	// godotenv
	_ = godotenv.Load()
	fmt.Println("Env file loaded (if exists)")

	// Example error
	err := errors.New("sample error")
	fmt.Println("Error with pkg/errors:", err)

	// Initialize router
	router := server.NewRouter()

	fmt.Println("Starting REST API server...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
