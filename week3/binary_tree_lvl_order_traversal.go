func levelOrder(root *TreeNode) [][]int {
    
    var ans [][]int
    var queue []*TreeNode
    
    queue = append(queue, root)
    
    for len(queue) > 0 {
        
        var lvl []int
        qLen := len(queue)
        
        for i:=0; i<qLen; i++ {
            top := queue[0]
            queue = queue[1:]
            
            if top != nil {
                lvl = append(lvl, top.Val)
                queue = append(queue, top.Left)
                queue = append(queue, top.Right)
            }
        }
        
        if len(lvl) > 0 {
            ans = append(ans,lvl)
        }
    }
    

    return ans
}
