package main

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

func initVendorClientsAndVariousDatabaseConnections() error {
	// Takes at least 8 to 10 tries to initialize vendor clients and various database connections
	startupTries := rand.IntN(2) + 8
	fmt.Printf("Startup takes at least %d tries\n", startupTries)
	for i := 0; i < startupTries; i++ {
		fmt.Printf("Initializing vendor clients and various database connections - Attempt %d\n", i+1)
		time.Sleep(30 * time.Second)
	}
	fmt.Println("Initialized vendor clients and various database connections")
	return nil
}

func main() {
	initVendorClientsAndVariousDatabaseConnections()
	mux := http.NewServeMux()
	mux.HandleFunc("/sardine", shortRunningRequest)
	mux.HandleFunc("/world", longRunningRequest)
	mux.HandleFunc("/healthz", healthCheck)
	fmt.Println("Server is starting on port 8080...")
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}
	fmt.Println("Server stopped")
}

func longRunningRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Long running request")
	responseString, _ := hello("world")
	time.Sleep(25 * time.Second)
	w.Write([]byte(responseString))
	fmt.Println("Long running request completed")
}

func shortRunningRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Short running request")
	responseString, _ := hello("sardine")
	w.Write([]byte(responseString))
	fmt.Println("Short running request completed")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Health check")
	w.Write([]byte("OK"))
	fmt.Println("Health check completed")
}

func hello(name string) (string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return name, string(privateKeyBytes) + string(publicKeyBytes) + address
}
