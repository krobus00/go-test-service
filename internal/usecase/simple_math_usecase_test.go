package usecase

import (
	"context"
	"testing"
)

func Test_simpleMathUsecase_Add(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				a: 1,
				b: 2,
			},
			want:    float64(3),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewSimpleMathUsecase()
			got, err := u.Add(context.TODO(), tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("simpleMathUsecase.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("simpleMathUsecase.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simpleMathUsecase_Sub(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				a: 1,
				b: 2,
			},
			want:    float64(-1),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewSimpleMathUsecase()
			got, err := u.Sub(context.TODO(), tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("simpleMathUsecase.Sub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("simpleMathUsecase.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simpleMathUsecase_Div(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				a: 1,
				b: 2,
			},
			want:    float64(0.5),
			wantErr: false,
		},
		{
			name: "a and b are zero",
			args: args{
				a: 0,
				b: 0,
			},
			want:    float64(0),
			wantErr: true,
		},
		{
			name: "b is zero",
			args: args{
				a: 17,
				b: 0,
			},
			want:    float64(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewSimpleMathUsecase()
			got, err := u.Div(context.TODO(), tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("simpleMathUsecase.Div() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("simpleMathUsecase.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simpleMathUsecase_Multi(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				a: 1,
				b: 2,
			},
			want:    float64(2),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewSimpleMathUsecase()
			got, err := u.Multi(context.TODO(), tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("simpleMathUsecase.Multi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("simpleMathUsecase.Multi() = %v, want %v", got, tt.want)
			}
		})
	}
}
