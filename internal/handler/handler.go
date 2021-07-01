package handler

import (
	"calc-api/internal/actions"
	"calc-api/internal/helper"
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
		// log.Println(err)
		_ = render.Render(w, r, helper.ErrBadRequest)
		return
	}

	num := &actions.Numbers{First: first, Second: second}
	result := num.Add()
	_ = render.Render(w, r, &CalcResponse{HTTPStatusCode: 200, Success: true, ErrCode: "", Value: result})
}

// вычитание
func Sub(w http.ResponseWriter, r *http.Request) {
	first, second, err := helper.GetParams(r)
	if err != nil {
		// log.Println(err)
		_ = render.Render(w, r, helper.ErrBadRequest)
		return
	}

	num := &actions.Numbers{First: first, Second: second}
	result := num.Sub()
	_ = render.Render(w, r, &CalcResponse{HTTPStatusCode: 200, Success: true, ErrCode: "", Value: result})
}

// умножение
func Mul(w http.ResponseWriter, r *http.Request) {
	first, second, err := helper.GetParams(r)
	if err != nil {
		// log.Println(err)
		_ = render.Render(w, r, helper.ErrBadRequest)
		return
	}

	num := &actions.Numbers{First: first, Second: second}
	result := num.Mul()
	_ = render.Render(w, r, &CalcResponse{HTTPStatusCode: 200, Success: true, ErrCode: "", Value: result})
}

// деление
func Div(w http.ResponseWriter, r *http.Request) {
	first, second, err := helper.GetParams(r)
	if err != nil {
		// log.Println(err)
		_ = render.Render(w, r, helper.ErrBadRequest)
		return
	}

	num := &actions.Numbers{First: first, Second: second}
	result, err := num.Div()
	if err != nil {
		_ = render.Render(w, r, helper.ErrBadRequest)
		return
	}

	_ = render.Render(w, r, &CalcResponse{HTTPStatusCode: 200, Success: true, ErrCode: "", Value: result})
}
