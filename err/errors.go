package err

import (
	"github.com/go-chi/render"
	"net/http"
)

const (
	ErrBadRequest       = "Некорректный запрос"
	ErrNotFound         = "Пользователь не найден"
	InternalServerError = "Внутренняя ошибка сервера"
)

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	IsBusinessErr bool   `json:"is_business_err"`
	StatusText    string `json:"status"`
	AppCode       int64  `json:"code,omitempty"`
	ErrorText     string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ResponseErr(statusCode int, isBusinessErr bool, statusText string, err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: statusCode,
		IsBusinessErr:  isBusinessErr,
		StatusText:     statusText,
		ErrorText:      err.Error(),
	}
}
