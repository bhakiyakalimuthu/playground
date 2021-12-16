package main

import (
	"testing"
)

func TestLRU(t *testing.T) {

	tests := map[string]struct {
		pairs    []pair
		capacity int
		want     []string
	}{

		"Should add and fetch key": {
			pairs:    []pair{{"key1", "value1"}, {"key2", "value2"}},
			capacity: 2,
			want:     []string{"value1", "value2"},
		},
		"Should add when cache is full": {
			pairs:    []pair{{"key1", "value1"}, {"key2", "value2"}, {"key3", "value3"}, {"key4", "value4"}, {"key5", "value5"}},
			capacity: 4,
			want:     []string{"value1", "value2", "value3", "value4", "value5"},
		},
		//"":{},

	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			c := New(testCase.capacity)
			for i := 0; i < len(testCase.pairs); i++ {
				c.Add(testCase.pairs[i].key, testCase.want[i])
				if got := c.Get(testCase.pairs[i].key); got != testCase.want[i] {
					t.Errorf("got %v not equal to want %v", got, testCase.want[i])
				}

			}

		})

	}
}
