package main

import "C"
import "fmt"

// var str string
// var str = ""

func main() {
	// [1,2,3,2,1]
	// [1,3,2,1,4,7]

	ret := sortedArrayToBST([]int{4, 2, 7, 1, 3}, 2)

	fmt.Println(ret)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if nil == root {
		return nil
	}
	if root.Val == val {
		return root
	}
	if lTree := searchBST(root.Left, val); nil != lTree {
		return lTree
	}
	return searchBST(root.Right, val)
}
