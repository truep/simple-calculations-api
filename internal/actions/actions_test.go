package actions

import (
	"testing"
)

func TestNumbers_Add(t *testing.T) {
	type fields struct {
		First  float64
		Second float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "Simple add test 1",
			fields: fields{1, 2},
			want:   3,
		},
		{
			name:   "Simple add test 2",
			fields: fields{10.24, 2.40},
			want:   12.64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Numbers{
				First:  tt.fields.First,
				Second: tt.fields.Second,
			}
			if got := n.Add(); got != tt.want {
				t.Errorf("Numbers.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumbers_Sub(t *testing.T) {
	type fields struct {
		First  float64
		Second float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "Simple sub test 1",
			fields: fields{1, 2},
			want:   -1,
		},
		{
			name:   "Simple sub test 2",
			fields: fields{10.24, 2.40},
			want:   7.84,
		},
		{
			name:   "Simple sub test 3",
			fields: fields{10, 2},
			want:   8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Numbers{
				First:  tt.fields.First,
				Second: tt.fields.Second,
			}
			if got := n.Sub(); got != tt.want {
				t.Errorf("Numbers.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumbers_Mul(t *testing.T) {
	type fields struct {
		First  float64
		Second float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "Simple mul test 1",
			fields: fields{2, 2},
			want:   4,
		},
		{
			name:   "Simple mul test 2",
			fields: fields{10.24, 2.40},
			want:   24.576,
		},
		{
			name:   "Simple mul test 3",
			fields: fields{10.24, 0},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Numbers{
				First:  tt.fields.First,
				Second: tt.fields.Second,
			}
			if got := n.Mul(); got != tt.want {
				t.Errorf("Numbers.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumbers_Div(t *testing.T) {
	type fields struct {
		First  float64
		Second float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name:    "Simple mul test 1",
			fields:  fields{10, 2},
			want:    5,
			wantErr: false,
		},
		{
			name:    "Simple mul test 2",
			fields:  fields{10.50, 2},
			want:    5.25,
			wantErr: false,
		},
		{
			name:    "Simple mul test 2",
			fields:  fields{10.50, 0},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Numbers{
				First:  tt.fields.First,
				Second: tt.fields.Second,
			}
			got, err := n.Div()
			if (err != nil) != tt.wantErr {
				t.Errorf("Numbers.Div() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Numbers.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}
