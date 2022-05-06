func climbStairs(n int) int {
    
    one,two := 1,1
    
    for i:=0; i<n; i++ {
        one,two = one+two, one
    }
    
    return two
}

