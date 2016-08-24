/*
	二叉查找树，也称有序二叉树(ordered binary tree)或已排序二叉树(sorted binary tree),
	是指一颗空树或者具有下列性质的二叉树:
	1. 若任意结点的左子树不空, 则左子树上所有结点的值均小于它的根结点的值;
	2. 若任意结点的右子树不空, 则右子树上所有结点的值均大于它的根结点的值;
	3. 任意结点的左子树和右子树也分别为二叉查找树.
	在一棵二叉查找树中, 是不会有重复的结点.
*/

package main

import "fmt"

// 二叉树结构
type BinNode struct {
	value  int
	left   *BinNode
	right  *BinNode
	parent *BinNode
}

func InitSortedBinaryTree(list []int) *BinNode {
	if len(list) == 0 {
		return nil
	}
	root := &BinNode{list[0], nil, nil, nil}
	for i := 1; i < len(list); i++ {
		Insert(root, list[i])
	}
	return root
}

// 插入z到树中
func Insert(root *BinNode, value int) {
	if root == nil {
		return
	}
	if root.value == value {
		return
	}
	z := &BinNode{value, nil, nil, nil}
	x := root
	y := x
	for x != nil {
		y = x
		if x.value > z.value {
			x = x.left
		} else if x.value < z.value {
			x = x.right
		} else {
			return
		}
	}
	z.parent = y
	if y.value > z.value {
		y.left = z
	} else {
		y.right = z
	}
	return
}

func main() {
	a := []int{4, 1, 2, 3, 6, 9}
	root := InitSortedBinaryTree(a)
	fmt.Printf("%v %v", a, root)
}
