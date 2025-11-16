package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"sct-backend-service/app/options"
)

func main() {
	// Parse command-line flags for local development
	port := flag.Int("port", 8080, "Server port (overrides config)")
	host := flag.String("host", "0.0.0.0", "Server host (overrides config)")
	debug := flag.Bool("debug", false, "Enable debug mode")
	configPath := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	log.Println("🚀 Starting local development server...")
	log.Printf("📝 Debug mode: %v", *debug)
	log.Printf("🌐 Host: %s", *host)
	log.Printf("🔌 Port: %d", *port)
	log.Printf("📁 Config: %s", *configPath)

	// Create fx application
	app := options.CreateApplication(*configPath)

	// TODO: Override config with command-line flags if needed

	// Check for errors in dependency graph
	if err := app.Err(); err != nil {
		log.Fatalf("Failed to create application: %v", err)
		os.Exit(1)
	}

	// Start the application (this triggers OnStart lifecycle hooks)
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatalf("Failed to start application: %v", err)
		os.Exit(1)
	}

	log.Println("✅ Application started successfully")
	log.Println("🌐 GraphQL endpoint: http://localhost:8080/query")
	log.Println("🎮 Playground: http://localhost:8080/")
	log.Println("⏹️  Press Ctrl+C to stop the server")

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Println("🛑 Shutting down server...")

	// Stop the application (this triggers OnStop lifecycle hooks)
	stopCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Printf("Error stopping application: %v", err)
	}

	log.Println("✅ Application shutdown complete")
}
