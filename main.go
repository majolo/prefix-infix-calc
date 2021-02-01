package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/akrylysov/algnhsa"
)

type Response struct {
	Result string `json:"result"`
	Error string `json:"error"`
}

type Request struct {
	Calculation string `json:"calculation"`
}

func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func prefixHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		// handle CORS pre-flight request
		setupResponse(&w)
		w.Write(nil)
		return
	}
	fmt.Println("received request for prefix")

	var req Request
	_ = json.NewDecoder(r.Body).Decode(&req)
	res, err := prefixCalculator(req.Calculation)
	retErr := ""
	if err != nil {
		// Would return 400 here if doing this in production.
		retErr = err.Error()
	}
	resp := Response{
		Result: fmt.Sprintf("%v", res),
		Error:  retErr,
	}
	j, _ := json.Marshal(resp)
	setupResponse(&w)
	w.Write(j)
}

func infixHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		// handle CORS pre-flight request
		setupResponse(&w)
		w.Write(nil)
		return
	}
	fmt.Println("received request for infix")

	var req Request
	_ = json.NewDecoder(r.Body).Decode(&req)
	res, err := infixCalculator(req.Calculation)
	retErr := ""
	if err != nil {
		// Would return 400 here if doing this in production.
		retErr = err.Error()
	}
	resp := Response{
		Result: fmt.Sprintf("%v", res),
		Error:  retErr,
	}
	j, _ := json.Marshal(resp)
	setupResponse(&w)
	w.Write(j)
}

func main() {
	// Setup handlers
	http.HandleFunc("/api/kheiron/prefix", prefixHandler)
	http.HandleFunc("/api/kheiron/infix", infixHandler)

	// If we're running in AWS or locally to make development easy
	if envVarExists("AWS_REGION") {
		fmt.Println("serving via lambda")
		algnhsa.ListenAndServe(http.DefaultServeMux, nil)

	} else {
		fmt.Println("serving on port 9000")
		_ = http.ListenAndServe(":9000", http.DefaultServeMux)
	}
}

func envVarExists(s string) bool {
	if _, ok := os.LookupEnv(s); ok {
		return true
	}
	return false
}