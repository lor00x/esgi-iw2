package math

import (
	"errors"
	"testing"
)

func TestAddPositive(t *testing.T) {
	t.Parallel()

	result := Add(1, 2)

	if result != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}
}

func TestAddNegative(t *testing.T) {
	t.Parallel()
	result := Add(-1, -2)

	if result != -3 {
		t.Errorf("Add(-1, -2) = %d; want -3", result)
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Basic: 1 + 1",
			args: args{a: 1, b: 1},
			want: 2,
		},
		{
			name: "Difficile: 1000 + 1",
			args: args{a: 1000, b: 1},
			want: 1001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr error
	}{
		{
			name: "Divide 10 / 2",
			args: args{
				a: 10,
				b: 2,
			},
			want:    5,
			wantErr: nil,
		},
		{
			name: "Divide by zero",
			args: args{
				a: 10,
				b: 0,
			},
			want:    0,
			wantErr: DivideByZero,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := Divide(tt.args.a, tt.args.b)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

			if got != tt.want {
				t.Errorf("Divide() got = %v, want %v", got, tt.want)
			}
		})
	}
}
