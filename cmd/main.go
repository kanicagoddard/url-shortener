package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"url-shortener/internal/handlers"
	"url-shortener/store/dbstore"

	"github.com/go-chi/chi/v5"
)

// using in-memory database
func main() {
	r := chi.NewRouter()

	// create an instace of a logger, using go's in-built structured logger, slog,introduced in 1.21
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// create instace of context
	// overall instance of context, all of the other instances of context will use this as base for their instance of context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// channel for kill signal
	killSig := make(chan os.Signal, 1)
	signal.Notify(killSig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// create instance of server
	svr := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	shortURLstore := dbstore.NewShortURLStore(dbstore.NewShortURLStoreParams{
		Logger: logger,
	})
	// graceful shutdown using go routine
	go func() {
		// wait for kill signal
		sig := <-killSig
		logger.Info("got kill signal >> shutting down!", slog.String("signal", sig.String()))

		// use timer to shutdown or else exit immediately with an error
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 5*time.Second)

		go func() {
			<-shutdownCtx.Done()

			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("shutdown deadline exceeded")
				os.Exit(1)
			}
		}()
		// Simulate delay in shutdown
		logger.Info("Simulating shutdown delay")
		time.Sleep(1 * time.Second)

		err := svr.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}

		serverStopCtx()
		logger.Info("server shutting down")
		cancel()
	}()

	go func() {
		err := svr.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

	}()
	r.Get("/healthcheck", handlers.NewHealthHandler().ServeHTTP)

	r.Post("/shorturl", handlers.NewCreateShortURLHandler(handlers.CreateShortURLHandlerParams{
		ShortURLStore: shortURLstore,
	}).ServerHTTP)

	logger.Info("ready to 🚀")

	<-serverCtx.Done()
}
