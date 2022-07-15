func insert(intervals [][]int, newInterval []int) [][]int {
    
    var res [][]int
    
    for i,interval := range intervals {
        if newInterval[1] < interval[0] {
            res = append(res,newInterval)
            res = append(res,intervals[i:]...)
            return res
        } else if newInterval[0] > interval[1] {
            res = append(res,interval)
        } else {
            newInterval[0] = min(newInterval[0], interval[0])
            newInterval[1] = max(newInterval[1], interval[1])
        }
    }
    
    res = append(res,newInterval)
    
    return res
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
