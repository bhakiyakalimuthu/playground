package testTable

import (
	"testing"
)

func TestProcess(t *testing.T) {
	var s Product
	testSuite := []struct {
		name          string
		field1        int
		expctedResult string
		expectErr     error
	}{
		{
			name:          "Test1",
			field1:        42,
			expctedResult: "Success",
			expectErr:     nil,
		},
		{
			name:          "Test2",
			field1:        -42,
			expctedResult: "Failure",
			expectErr:     nil,
		},
	}

	for i := 0; i < 100; i++ {
		for _, tt := range testSuite {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				s.Process(tt.field1)
			})
		}
	}
}
