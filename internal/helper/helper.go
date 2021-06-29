package helper

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

/*{
    "Success":true,
    "ErrCode":"",
    "Value":16
}*/
type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	Success bool    `json:"Success"`
	ErrCode string  `json:"ErrCode"`
	Value   float64 `json:"Value,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, Success: false, ErrCode: "404"}
var ErrBadRequest = &ErrResponse{HTTPStatusCode: 400, Success: false, ErrCode: "400"}

func GetParams(r *http.Request) (float64, float64, error) {
	var a, b string
	if a = r.URL.Query().Get("a"); a == "" {
		return 0, 0, errors.New("There is no 'a' parameter at query")
	}
	if b = r.URL.Query().Get("b"); b == "" {
		return 0, 0, errors.New("There is no 'b' parameter at query")
	}

	first, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, 0, err
	}
	second, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, 0, err
	}

	return first, second, nil
}
