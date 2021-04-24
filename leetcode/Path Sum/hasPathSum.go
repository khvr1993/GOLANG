package hasPathSum

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return addSumalongLeaf(root, 0, targetSum)
}

func addSumalongLeaf(node *TreeNode, currentSum int, targetSum int) bool {

	var currSum int = currentSum

	//exit condition
	if node == nil {
		return false
	}
	currSum += node.Val

	if currSum == targetSum && node.Left == nil && node.Right == nil {
		return true
	} else {
		return addSumalongLeaf(node.Left, currSum, targetSum) || addSumalongLeaf(node.Right, currSum, targetSum)
	}

}
