func maxSubArray(nums []int) int {
    maxTillNow := -100000
    currMax := 0
    
    for _,num := range nums {
        currMax = max(num, num+currMax)
        if currMax>maxTillNow {
            maxTillNow = currMax
        }
    }
    
    return maxTillNow
}

func max(x int,y int) int{
    if x>y {
        return x
    }
    return y
}
