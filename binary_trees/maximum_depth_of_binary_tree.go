package binary_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := root.Left
	right := root.Right

	return max(MaxDepth(left), MaxDepth(right)) + 1

}
