package queue

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want *Queue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
