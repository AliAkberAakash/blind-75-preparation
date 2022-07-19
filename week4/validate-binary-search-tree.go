func isValidBST(root *TreeNode) bool {
    
    const MaxUint = ^uint(0) 
    const MinUint = 0 
    const MaxInt = int(MaxUint >> 1) 
    const MinInt = -MaxInt - 1
    
    return validate(root,MinInt,MaxInt)
}

func validate(root *TreeNode, left, right int) bool {
    if root == nil {
        return true
    }
    
    if !(root.Val < right && root.Val > left) {
        return false
    }

    
    return validate(root.Left, left, root.Val) && validate(root.Right, root.Val, right)
}
