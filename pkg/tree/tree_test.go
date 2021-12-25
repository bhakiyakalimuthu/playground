package tree

import (
	"reflect"
	"testing"
)

func Test_treeInOrderTraversal(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"success case 1": {
			[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5},
		},
		"success case 2": {
			[]int{70, 60, 80, 90, 50}, []int{50, 60, 70, 80, 90},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			tree := treeInit(testCase.input)
			if got := tree.InOrder(); reflect.DeepEqual(got, testCase.want) {
				t.Logf("got : %v = want : %v", got, testCase.want)
			} else {
				t.Errorf("got : %v != want : %v", got, testCase.want)
			}
		})
	}
}
func Test_treePreOrderTraversal(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"success case 1": {
			[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5},
		},
		"success case 2": {
			[]int{70, 60, 80, 90, 50}, []int{50, 60, 70, 80, 90},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			tree := treeInit(testCase.input)
			if got := tree.PreOrder(); reflect.DeepEqual(got, testCase.want) {
				t.Logf("got : %v = want : %v", got, testCase.want)
			} else {
				t.Errorf("got : %v != want : %v", got, testCase.want)
			}
		})
	}
}
func Test_treePostOrderTraversal(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"success case 1": {
			[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5},
		},
		"success case 2": {
			[]int{70, 60, 80, 90, 50}, []int{50, 60, 70, 80, 90},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			tree := treeInit(testCase.input)
			if got := tree.PostOrder(); reflect.DeepEqual(got, testCase.want) {
				t.Logf("got : %v = want : %v", got, testCase.want)
			} else {
				t.Errorf("got : %v != want : %v", got, testCase.want)
			}
		})
	}
}
func treeInit(in []int) *BinaryTree {
	tree := New()
	for _, val := range in {
		tree.Insert(val)
	}
	return tree
}
