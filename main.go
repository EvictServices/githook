package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
	"github.com/shi-gg/githook/config"
	"github.com/shi-gg/githook/routes"
)

var conf = config.Get()

var client = redis.NewClient(&redis.Options{
	Addr:     conf.Redis.Addr,
	Password: conf.Redis.Password,
	Username: conf.Redis.Username,
	DB:       conf.Redis.Db,
})

func main() {
	defer client.Close()

	http.HandleFunc("POST /incoming/{id}", func(w http.ResponseWriter, r *http.Request) {
		routes.HandleIncoming(w, r, client)
	})

	http.HandleFunc("GET /create", routes.HandleCreate)

	server := &http.Server{Addr: ":8080"}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v\n", err)
		}
	}()

	<-stop
	log.Println("Shutting down...")
	server.Shutdown(context.Background())
}
