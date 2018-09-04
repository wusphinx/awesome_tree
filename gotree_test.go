package gotree_test

import (
	"testing"

	"github.com/wusphinx/gotree"
)

/*
		 4
		/ \
	   3   8
	  /     \
	 1       9
*/

func TestInitSortedBinaryTree(t *testing.T) {
	tmp := []int{4, 3, 1, 8, 9}
	root := gotree.InitSortedBinaryTree(tmp)
	if root.Value == 4 {
		a := root.Left
		b := a.Left
		c := root.Right
		d := c.Right
		if !(a.Value == 3 && b.Value == 1 && c.Value == 8 && d.Value == 9) {
			t.Errorf("︶︿︶ tree init failed")
		} else {
			t.Logf("^_^ tree init ok")
		}
	} else {
		t.Errorf("︶︿︶ root value should be %v", tmp[0])
	}
}
