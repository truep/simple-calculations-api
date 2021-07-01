package actions

import (
	"errors"
	"math"
)

type Numbers struct {
	First  float64
	Second float64
}

func (n *Numbers) Add() float64 {
	return n.First + n.Second
}

func (n *Numbers) Sub() float64 {
	return n.First - n.Second
}
func (n *Numbers) Mul() float64 {
	return n.First * n.Second
}
func (n *Numbers) Div() (float64, error) {
	result := n.First / n.Second

	// zero division check
	if math.IsInf(result, 0) {
		return 0, errors.New("Division error")
	} else {
		return result, nil
	}
}
