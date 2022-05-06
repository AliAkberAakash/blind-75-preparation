/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    leftHeight := treeHeight(root.Left)
    rightHeight := treeHeight(root.Right)
    
    if abs(leftHeight-rightHeight) > 1 {
        return false
    }
    
    return isBalanced(root.Left) && isBalanced(root.Right)
}

func treeHeight(root *TreeNode) int {
    if root == nil {
        return 1
    }
    
    left := 1+treeHeight(root.Left)
    right := 1+treeHeight(root.Right)
    
    return max(left,right)
}

func max(x int, y int) int {
    if x>y {
        return x
    }
    return y
}

func abs(x int) int {
    if x<0 {
        return -x
    }
    return x
}
