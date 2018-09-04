/*
	二叉查找树，也称有序二叉树(ordered binary tree)或已排序二叉树(sorted binary tree),
	是指一颗空树或者具有下列性质的二叉树:
	1. 若任意结点的左子树不空, 则左子树上所有结点的值均小于它的根结点的值;
	2. 若任意结点的右子树不空, 则右子树上所有结点的值均大于它的根结点的值;
	3. 任意结点的左子树和右子树也分别为二叉查找树.
	在一棵二叉查找树中, 是不会有重复的结点.
*/

package gotree

// 二叉树结构
type BinNode struct {
	Value  int
	Left   *BinNode
	Right  *BinNode
	Parent *BinNode
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
func Insert(root *BinNode, Value int) {
	if root == nil {
		return
	}
	if root.Value == Value {
		return
	}
	z := &BinNode{Value, nil, nil, nil}
	x := root
	y := x
	for x != nil {
		y = x
		if x.Value > z.Value {
			x = x.Left
		} else if x.Value < z.Value {
			x = x.Right
		} else {
			return
		}
	}
	z.Parent = y
	if y.Value > z.Value {
		y.Left = z
	} else {
		y.Right = z
	}
	return
}
