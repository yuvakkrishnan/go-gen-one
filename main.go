package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/gorilla/mux"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println("SBOM Test Project with Random Go Libraries")

	// Using uuid
	id := uuid.New()
	fmt.Println("UUID:", id)

	// Using logrus
	log := logrus.New()
	log.Info("Logrus initialized")

	// Cobra (CLI framework)
	var rootCmd = &cobra.Command{Use: "sbomtest"}
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("Cobra CLI executed")
	}
	_ = rootCmd.Execute()

	// Viper (config)
	viper.SetDefault("port", 8080)
	fmt.Println("Default port:", viper.GetInt("port"))

	// Gorilla mux
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from mux router")
	})
	fmt.Println("Router initialized")

	// Resty HTTP client
	client := resty.New()
	_, _ = client.R().Get("https://httpbin.org/get")
	fmt.Println("Resty client executed GET")

	// godotenv
	_ = godotenv.Load()
	fmt.Println("Env file loaded (if exists)")

	// JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	ss, _ := token.SignedString([]byte("secret"))
	fmt.Println("JWT Token:", ss)

	// bcrypt
	hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	fmt.Println("Password hash:", string(hash))

	// pkg/errors
	err := errors.New("sample error")
	fmt.Println("Error with pkg/errors:", err)
}
