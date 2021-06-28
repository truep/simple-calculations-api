package handler

import "net/http"

/*
api methods: add, sub, mul, div

response ex.:
{
    "Success":true,
    "ErrCode":"",
    "Value":16
}

*/

func Add(w http.ResponseWriter, r *http.Request) {}
func Sub(w http.ResponseWriter, r *http.Request) {}
func Mul(w http.ResponseWriter, r *http.Request) {}
func Div(w http.ResponseWriter, r *http.Request) {}
