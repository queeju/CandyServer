package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type BuyCandyRequest struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

type BuyCandyResponse struct {
	Thanks string `json:"thanks"`
	Change int    `json:"change"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	candyType := flag.String("k", "", "Type of candy")
	candyCount := flag.Int("c", 0, "Count of candy")
	cowSay := flag.Bool("f", false, "Enable cowsay response")
	money := flag.Int("m", 0, "Amount of money")
	flag.Parse()

	caCert, err := os.ReadFile("./cert/ca-cert.pem")
	if err != nil {
		log.Fatalf("Failed to read CA cert: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Load client cert
	cert, err := tls.LoadX509KeyPair("./cert/client-cert.pem", "./cert/client-key.pem")
	if err != nil {
		log.Fatalf("Failed to load client cert/key: %v", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	buyCandyRequest := BuyCandyRequest{
		Money:      *money,
		CandyType:  *candyType,
		CandyCount: *candyCount,
	}

	requestBody, err := json.Marshal(buyCandyRequest)
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v", err)
	}

	resp, err := client.Post("https://127.0.0.1:3333/buy_candy", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusCreated {
		var response BuyCandyResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			log.Fatalf("Failed to decode response: %v", err)
		}
		if *cowSay {
			fmt.Printf("Your change is %d\n", response.Change)
			fmt.Printf("%s", response.Thanks)
		} else {
			fmt.Printf("Thank you! Your change is %d\n", response.Change)
		}
	} else {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			log.Fatalf("Failed to decode error response: %v", err)
		}
		fmt.Printf("Error: %s\n", errorResponse.Error)
	}
}
