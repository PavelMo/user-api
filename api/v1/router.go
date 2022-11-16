package v1

import (
	u "github.com/PavelMo/user-api/api/v1/use-cases"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

func AttachRouter(r *chi.Mux, u *u.UserUseCases) http.Handler {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(time.Now().String()))
		if err != nil {
			log.Println(err)
		}
	})

	r.Mount("/users", attachUsersRouter(r, u))
	return r
}

func attachUsersRouter(r *chi.Mux, u *u.UserUseCases) http.Handler {
	r.Get("/", u.SearchUsers)
	r.Post("/", u.CreateUser)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", u.GetUser)
		r.Patch("/", u.UpdateUser)
		r.Delete("/", u.DeleteUser)
	})

	return r
}
