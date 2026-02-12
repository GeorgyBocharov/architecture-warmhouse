package main

import (
	"fmt"

	"temperature/internal/service"
	"temperature/server"
	"net"
	"net/http"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	log.Printf("starting app")
	swagger, err := server.GetSwagger()
  	if err != nil {
    	fmt.Printf("error loading swagger spec\n: %s", err)

    	return
 	}

  // Clear out the servers array in the swagger spec, that skips validating
  // that server names match. We don't know how this thing will be run.
  swagger.Servers = nil

  router := chi.NewRouter()
	
  server.HandlerFromMux(server.NewTemperatureAdapter(&service.RandomTemperatureProcessor{}), router)

  srv := &http.Server{
		Addr:    net.JoinHostPort(getEnv("HOST", "0.0.0.0"), getEnv("PORT", "8080")),
		Handler: router,
	}

	// Start the server in a goroutine
	go func() {
		log.Printf("Server starting on %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server exited properly")
	
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}