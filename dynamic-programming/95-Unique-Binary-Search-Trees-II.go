package dynamic_programming

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return help95(1, n)
}

func help95(start, end int) []*TreeNode {
	var res []*TreeNode
	if start > end {
		res = append(res, nil)
		return res
	}

	for i := start; i <= end; i++ {
		left := help95(start, i-1)
		right := help95(i+1, end)
		for _, lnode := range left {
			for _, rnode := range right {
				node := &TreeNode{}
				node.Val = i
				node.Left = lnode
				node.Right = rnode
				res = append(res, node)
			}
		}
	}
	return res
}
