package handlers

import (
	"fmt"
	"net/http"

	"go-gen-one/utils"
)

// HomeHandler simple hello
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from mux router")
}

// JWTHandler returns a signed JWT token
func JWTHandler(w http.ResponseWriter, r *http.Request) {
	token, err := utils.GenerateJWT("secret")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "JWT Token:", token)
}

// HashHandler returns bcrypt hash
func HashHandler(w http.ResponseWriter, r *http.Request) {
	hash, _ := utils.GenerateHash("password")
	fmt.Fprintln(w, "Password hash:", hash)
}

// RestyHandler calls httpbin.org/get
func RestyHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := utils.DoRestyGet("https://httpbin.org/get")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Resty GET Response:", resp)
}
