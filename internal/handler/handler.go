package handler

import (
	"calc-api/internal/helper"
	"log"
	"math"
	"net/http"

	"github.com/go-chi/render"
)

/*{
    "Success":true,
    "ErrCode":"",
    "Value":16
}*/
type CalcResponse struct {
	HTTPStatusCode int `json:"-"`

	Success bool    `json:"Success"`
	ErrCode string  `json:"ErrCode"`
	Value   float64 `json:"Value"`
}

func (c *CalcResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, c.HTTPStatusCode)
	return nil
}

// сложение
func Add(w http.ResponseWriter, r *http.Request) {
	first, second, err := helper.GetParams(r)
	if err != nil {
		log.Println(err)
		_ = render.Render(w, r, helper.ErrBadRequest)
		return
	}

	result := first + second
	_ = render.Render(w, r, &CalcResponse{HTTPStatusCode: 200, Success: true, ErrCode: "200", Value: result})
}

// вычитание
func Sub(w http.ResponseWriter, r *http.Request) {
	first, second, err := helper.GetParams(r)
	if err != nil {
		log.Println(err)
		_ = render.Render(w, r, helper.ErrBadRequest)
		return
	}

	result := first - second
	_ = render.Render(w, r, &CalcResponse{HTTPStatusCode: 200, Success: true, ErrCode: "200", Value: result})
}

// умножение
func Mul(w http.ResponseWriter, r *http.Request) {
	first, second, err := helper.GetParams(r)
	if err != nil {
		log.Println(err)
		_ = render.Render(w, r, helper.ErrBadRequest)
		return
	}

	result := first * second
	_ = render.Render(w, r, &CalcResponse{HTTPStatusCode: 200, Success: true, ErrCode: "200", Value: result})
}

// деление
func Div(w http.ResponseWriter, r *http.Request) {
	first, second, err := helper.GetParams(r)
	if err != nil {
		log.Println(err)
		_ = render.Render(w, r, helper.ErrBadRequest)
		return
	}

	result := first / second

	// zero division check
	if math.IsInf(result, 0) {
		log.Println(result)
		_ = render.Render(w, r, helper.ErrBadRequest)
		return
	}

	_ = render.Render(w, r, &CalcResponse{HTTPStatusCode: 200, Success: true, ErrCode: "200", Value: result})
}
