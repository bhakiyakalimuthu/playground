package main

import (
	"sync"
	"testing"
)

func Test_deposit(t *testing.T) {
	type args struct {
		amount float32
		wg     *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add testplay cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
