package main

import "testing"

func TestInitSortedBinaryTree(t *testing.T) {
	tmp := []int{4, 3, 1, 8, 9}
	root := InitSortedBinaryTree(tmp)
	if root.value == 4 {
		a := root.left
		b := a.left
		c := root.right
		d := c.right
		if !(a.value == 3 && b.value == 1 && c.value == 8 && d.value == 9) {
			t.Errorf("︶︿︶ tree init failed")
		} else {
			t.Logf("^_^ tree init ok")
		}
	} else {
		t.Errorf("︶︿︶ root value should be %v", tmp[0])
	}
}
