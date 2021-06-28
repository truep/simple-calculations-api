package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/ilyakaznacheev/cleanenv"

	"calc-api/internal/config"
	"calc-api/internal/handler"
)

func main() {
	cfg := config.Server{}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal(err)
	}

	path := cfg.String()
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(30*time.Second),
	)

	// ex.:http://localhost/api/add?a=1&b=15
	r.Route("/api", func(r chi.Router) {
		// r.With()
		r.Get("/add", handler.Add)
		r.Get("/sub", handler.Sub)
		r.Get("/mul", handler.Mul)
		r.Get("/div", handler.Div)
	})

	srv := &http.Server{
		Addr:    path,
		Handler: r,
	}

	// shutdown gracefully
	quit := make(chan os.Signal, 1)
	done := make(chan error, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		err := srv.Shutdown(ctx)
		done <- err
	}()

	log.Printf("Starting server at %s", path)
	_ = srv.ListenAndServe()

	err := <-done

	log.Printf("Shutting server down with %v", err)
}
