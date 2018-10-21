package main

import (
	"os"
	"fmt"
	"time"
	"net/http"
)

import (
	"github.com/edinkulovic/BasicHTTPServerWithCORS/config"
	"github.com/edinkulovic/BasicHTTPServerWithCORS/routes"
)

// HTTPServer - Server handler
type HTTPServer struct{}

// Mux - Contains routes and handlers
var Mux map[string]http.HandlerFunc

func main() {
	fmt.Printf(" * Server starting at port: %v\n", config.Service.Port)

	// Initialize Database

	// Initialize Redis

	// Setup HTTP Server
	server := http.Server{
		Addr:				":" + config.Service.Port,
		Handler:			&HTTPServer{},
		ReadTimeout:       	config.Timeouts.Read * time.Second,
		WriteTimeout:      	config.Timeouts.Write * time.Second,
		ReadHeaderTimeout: 	config.Timeouts.ReadHeader * time.Second,
		IdleTimeout:       	config.Timeouts.Idle * time.Second,
	}

	Mux = routes.Mux 

	err := server.ListenAndServe()

	if err != nil {
		fmt.Printf("Unable to listen on HTTP Server: Error: %v", err)
		os.Exit(1)
	}
}

func (*HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "Token, Status-Code, Status-Text, Set-Cookie")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept-Encoding, Accept, Access-Control-Allow-Origin, Authorization, Token, Status-Code, Status-Text, Set-Cookie, Access-Control-Allow-Credentials")

	if r.Method == http.MethodOptions {
		return
	}

	if h, ok := Mux[r.URL.Path]; ok { // Check if route exists
		h(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	}
}
