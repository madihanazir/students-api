package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	//"github.com/gorilla/mux"
	"github.com/madihanazir/students-api/internal/config"
	"github.com/madihanazir/students-api/internal/http/handlers/student"
	"github.com/madihanazir/students-api/storage/sqlite"
)

func main() {
	//loag config & database setup
	cfg := config.Mustload()

	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatalf("can't connect to database: %v", err)
	}
	
	slog.Info("storage initialized", slog.String("version", "1.0.0"))

	//setup router
	router := http.NewServeMux()

	fmt.Println("Registered Routes:")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Root route working!"))})


	fmt.Println("/ registered")
	fmt.Println("/api/students registered")

	router.HandleFunc("/api/students", student.New(storage))
	fmt.Println("/api/students registered")

	

	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test route working!")
	})
	

	//router.HandleFunc("/api/students", student.New().ServeHTTP)
	//http.HandleFunc("/api/students", student.New())

	//setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}
	slog.Info("starting server...", slog.String("address", cfg.HTTPServer.Addr))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("can't start server %v", err)
		}
	}()
	<-done

	slog.Info("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("can't shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("server stopped")
	
}
