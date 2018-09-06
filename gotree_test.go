package gotree_test

import (
	"reflect"
	"testing"

	"github.com/wusphinx/gotree"
)

/*
		  4
		/   \
	   2    10
	  / \   /
	 1   3 5
			\
			 8
			/ \
		   7   9
		  /
		 6
*/

func TestInitSortedBinaryTree(t *testing.T) {
	testCases := []struct {
		input []int
		dfs   []int
		bfs   []int
	}{
		{
			input: []int{4, 2, 3, 10, 1, 5, 8, 9, 7, 6},
			dfs:   []int{4, 2, 1, 3, 10, 5, 8, 7, 6, 9},
			bfs:   []int{4, 2, 10, 1, 3, 5, 8, 7, 9, 6},
		},
		{
			input: []int{1},
			dfs:   []int{1},
			bfs:   []int{1},
		},
		{
			input: []int{1, 2, 3},
			dfs:   []int{1, 2, 3},
			bfs:   []int{1, 2, 3},
		},
		{
			input: []int{3, 2, 1},
			dfs:   []int{3, 2, 1},
			bfs:   []int{3, 2, 1},
		},
	}

	for _, tc := range testCases {
		root := gotree.InitSortedBinaryTree(tc.input)
		dfs := gotree.DepthFirstSearch(root)
		t.Logf("dfs:%+v", dfs)
		if !reflect.DeepEqual(dfs, tc.dfs) {
			t.FailNow()
		}

		bfs := gotree.BreadthFirstSearch(root)
		t.Logf("bfs:%+v", bfs)
		if !reflect.DeepEqual(bfs, tc.bfs) {
			t.FailNow()
		}
	}
}
