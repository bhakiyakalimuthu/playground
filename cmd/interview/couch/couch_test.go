package main

import (
	"testing"
)

func TestConvert(t *testing.T) {

	tests := map[string]struct {
		data    interface{}
		want    string
		wantErr bool
	}{
		"valid string 1": {
			data:    "18.00",
			want:    "18.00",
			wantErr: false,
		},
		"valid string 2": {
			data:    "100000000.00",
			want:    "100000000.00",
			wantErr: false,
		},
		"invalid string ": {
			data:    "its invalid string",
			want:    "",
			wantErr: true,
		},
		"string with dollar in prefix": {
			data:    "$100.00",
			want:    "100.00",
			wantErr: false,
		},
		"string with comma": {
			data:    "$10,07,93,783748.00",
			want:    "100793783748.00",
			wantErr: false,
		},
		"string with dollar sign in the end": {
			data:    "100.00$",
			want:    "100.00",
			wantErr: false,
		},
		"nil interface": {
			data:    nil,
			want:    "",
			wantErr: true,
		},
		"int value": {
			data:    852898,
			want:    "852898.00",
			wantErr: false,
		},
		"int32 value": {
			data:    8048131,
			want:    "8048131.00",
			wantErr: false,
		},
		"int64 value": {
			data:    749327492,
			want:    "749327492.00",
			wantErr: false,
		},
		"float32 value": {
			data:    532615362.73827,
			want:    "532615360.00",
			wantErr: false,
		},
		"float64 value": {
			data:    4273932.89,
			want:    "4273933.00",
			wantErr: false,
		},
		"byte value": {
			data:    []byte(`457347`),
			want:    "457347.00",
			wantErr: false,
		},
		"byte value with $": {
			data:    []byte(`$8932`),
			want:    "8932.00",
			wantErr: false,
		},
		"byte value with not support": {
			data:    []byte(`life is beautiful`),
			want:    "",
			wantErr: true,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := convert(testCase.data)
			if (err != nil) != testCase.wantErr {
				t.Errorf("convert() error= %v, wantErr:%v", err, testCase.wantErr)
			}
			if got != testCase.want {
				t.Errorf("convert() got: %v, want: %v", testCase.want, got)
			}
		})
	}
}

func BenchmarkConvert(t *testing.B) {

	tests := map[string]struct {
		data    interface{}
		want    string
		wantErr bool
	}{
		"valid string 1": {
			data:    "18.00",
			want:    "18.00",
			wantErr: false,
		},
		"valid string 2": {
			data:    "100000000.00",
			want:    "100000000.00",
			wantErr: false,
		},
		"invalid string ": {
			data:    "its invalid string",
			want:    "",
			wantErr: true,
		},
		"string with dollar in prefix": {
			data:    "$100.00",
			want:    "100.00",
			wantErr: false,
		},
		"string with dollar sign in the end": {
			data:    "100.00$",
			want:    "100.00",
			wantErr: false,
		},
		"nil interface": {
			data:    nil,
			want:    "",
			wantErr: true,
		},
		"int value": {
			data:    852898,
			want:    "852898.00",
			wantErr: false,
		},
		"int32 value": {
			data:    8048131,
			want:    "8048131.00",
			wantErr: false,
		},
		"int64 value": {
			data:    749327492,
			want:    "749327492.00",
			wantErr: false,
		},
		"float32 value": {
			data:    532615362.73827,
			want:    "532615360.00",
			wantErr: false,
		},
		"float64 value": {
			data:    4273932.89,
			want:    "4273933.00",
			wantErr: false,
		},
		"byte value": {
			data:    []byte(`457347`),
			want:    "457347.00",
			wantErr: false,
		},
		"byte value with $": {
			data:    []byte(`$8932`),
			want:    "8932.00",
			wantErr: false,
		},
		"byte value with not support": {
			data:    []byte(`life is beautiful`),
			want:    "",
			wantErr: true,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.B) {
			got, err := convert(testCase.data)
			if (err != nil) != testCase.wantErr {
				t.Errorf("convert() error= %v, wantErr:%v", err, testCase.wantErr)
			}
			if got != testCase.want {
				t.Errorf("convert() got: %v, want: %v", testCase.want, got)
			}
		})
	}
}
