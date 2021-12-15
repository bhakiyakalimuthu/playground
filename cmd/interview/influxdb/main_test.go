package main

import "testing"

func TestIsPresent(t *testing.T){

	tests := map[string]struct{

		in int64
		want bool
	}{
		"case 1": {
			in: 1639499933,
			want:true,
		},
		"case 2": {
			in: 1639586009,
			want:true,
		},
		"case 3": {
			in: 1639508264,
			want:true,
		},

	}

	for testName,testCase := range tests {
		shards := []shard{
			{1639495472,1639498974},
			{1639491888,1639506294},
			{1639502701,1639499726},
		}
		t.Run(testName,func(t *testing.T){
			if got := isPointPresent(shards,testCase.in); got != testCase.want {
				t.Errorf("want %v != got %v", testCase.want,got)
			}
		})


	}

}
