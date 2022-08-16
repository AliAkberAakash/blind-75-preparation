func combinationSum(candidates []int, target int) [][]int {
 
    limit := len(candidates)
    var result [][]int
    var current []int 
    
    dfs(0,current,0,target,limit, candidates, &result)
    
    return result
}

func dfs(i int, current []int, total int, target int, limit int, candidates []int, result* [][]int) {
    
    if total == target {
        tmp := make([]int, len(current))
        copy(tmp,current)
        
        (*result) = append(*result, tmp)
        return
    }
    
    if i >= limit || total > target {
        return
    }
    
    current = append(current,candidates[i])
    dfs(i, current, total+candidates[i], target, limit, candidates, result)
    current = current[:len(current)-1]
    dfs(i+1, current, total, target, limit, candidates, result)
    
}
