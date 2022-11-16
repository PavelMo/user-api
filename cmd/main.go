package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	v1 "refactoring/api/v1"
	u "refactoring/api/v1/use-cases"
	"refactoring/impl/repo"
	"refactoring/impl/services"
	"syscall"
	"time"
)

func main() {
	userApi := u.NewUserUseCases(services.NewUserService(repo.NewUserRepo()))
	r := chi.NewRouter()
	r.Mount("/api/v1", v1.AttachRouter(r, userApi))

	srv := &http.Server{
		Addr:    ":3333",
		Handler: r,
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		log.Printf("service started at port %s \n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	<-sig
	log.Println("shutting down service...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}

}
