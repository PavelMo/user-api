package use_cases

import (
	"errors"
	"github.com/PavelMo/user-api/api/v1/views"
	errs "github.com/PavelMo/user-api/err"
	"github.com/PavelMo/user-api/interfaces/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type UserUseCases struct {
	srv services.IUserService
}

func NewUserUseCases(service services.IUserService) *UserUseCases {
	return &UserUseCases{srv: service}
}

func (u *UserUseCases) SearchUsers(w http.ResponseWriter, r *http.Request) {
	res := u.srv.GetAllUsers()

	render.JSON(w, r, res)
}

func (u *UserUseCases) CreateUser(w http.ResponseWriter, r *http.Request) {
	request := views.CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, errs.ResponseErr(http.StatusBadRequest, false, errs.ErrBadRequest, errors.New(errs.ErrBadRequest)))

		return
	}
	id, err := u.srv.CreateUser(request.DisplayName, request.Email)
	if err != nil {
		_ = render.Render(w, r, errs.ResponseErr(http.StatusInternalServerError, true, errs.InternalServerError, errors.New(errs.InternalServerError)))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func (u *UserUseCases) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		_ = render.Render(w, r, errs.ResponseErr(http.StatusBadRequest, false, errs.ErrBadRequest, errors.New(errs.ErrBadRequest)))
		return
	}
	user, err := u.srv.GetUser(id)
	if err != nil {
		_ = render.Render(w, r, errs.ResponseErr(http.StatusNotFound, false, errs.ErrNotFound, errors.New(errs.ErrNotFound)))
		return
	}

	render.JSON(w, r, user)
}

func (u *UserUseCases) UpdateUser(w http.ResponseWriter, r *http.Request) {
	request := views.UpdateUserRequest{}
	id := chi.URLParam(r, "id")

	if err := render.Bind(r, &request); err != nil || id == "" {
		_ = render.Render(w, r, errs.ResponseErr(http.StatusBadRequest, false, errs.ErrBadRequest, errors.New(errs.ErrBadRequest)))
		return
	}

	if err := u.srv.UpdateUser(id, request.DisplayName); err != nil {
		if err.Error() == errs.ErrNotFound {
			_ = render.Render(w, r, errs.ResponseErr(http.StatusNotFound, false, errs.ErrNotFound, errors.New(errs.ErrNotFound)))
		} else {
			_ = render.Render(w, r, errs.ResponseErr(http.StatusInternalServerError, true, errs.InternalServerError, errors.New(errs.InternalServerError)))
		}

		return
	}

	render.Status(r, http.StatusNoContent)
}

func (u *UserUseCases) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		_ = render.Render(w, r, errs.ResponseErr(http.StatusBadRequest, false, errs.ErrBadRequest, errors.New(errs.ErrBadRequest)))
		return
	}

	if err := u.srv.DeleteUser(id); err != nil {
		if err.Error() == errs.ErrNotFound {
			_ = render.Render(w, r, errs.ResponseErr(http.StatusNotFound, false, errs.ErrNotFound, errors.New(errs.ErrNotFound)))
		} else {
			_ = render.Render(w, r, errs.ResponseErr(http.StatusInternalServerError, true, errs.InternalServerError, errors.New(errs.InternalServerError)))
		}

		return
	}

	render.Status(r, http.StatusNoContent)
}
