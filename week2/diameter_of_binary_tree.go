/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    
    max := 0
    
    diameter(root,&max)
    
    return max
}

func diameter(root *TreeNode, max *int) int {
    if root == nil {
        return -1
    }
    
    left := diameter(root.Left, max)
    right := diameter(root.Right, max)
    
    *max = myMax(*max, 2 + left + right)
    
    return 1 + myMax(left,right)
}

func myMax(a,b int) int {
    if a>b {
        return a
    }
    return b
}
