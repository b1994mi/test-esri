package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/uptrace/bunrouter"
)

func main() {
	

	r := bunrouter.New()
	r.GET("/", func(w http.ResponseWriter, req bunrouter.Request) error {
		bunrouter.JSON(w, bunrouter.H{
			"message": "pong",
		})
		return nil
	})

	r.GET("/user", func(w http.ResponseWriter, req bunrouter.Request) error {
		bunrouter.JSON(w, bunrouter.H{
			"data": nil,
		})
		return nil
	})

	port := ":5000"
	httpLn, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	httpServer := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      r,
	}

	go func() {
		log.Printf("running on port %v", port)
		err := httpServer.Serve(httpLn)
		if err != nil {
			log.Println(err)
		}
	}()

	log.Println(waitExitSignal().String())

	ctx := context.Background()
	// Graceful shutdown.
	err = httpServer.Shutdown(ctx)
	if err != nil {
		log.Printf("trying to shutdown with error: %v", err)
	}
}

func waitExitSignal() os.Signal {
	ch := make(chan os.Signal, 3)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	return <-ch
}
