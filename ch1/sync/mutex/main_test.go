package main

import (
	"sync"
	"testing"
)

func TestCounter_Increment(t *testing.T) {
	type fields struct {
		Mutex sync.Mutex
		value int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Counter{
				Mutex: tt.fields.Mutex,
				value: tt.fields.value,
			}
			c.Increment()
		})
	}
}

func TestCounter_Value(t *testing.T) {
	type fields struct {
		Mutex sync.Mutex
		value int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Counter{
				Mutex: tt.fields.Mutex,
				value: tt.fields.value,
			}
			if got := c.Value(); got != tt.want {
				t.Errorf("Counter.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
