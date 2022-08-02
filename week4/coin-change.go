func coinChange(coins []int, amount int) int {
    
    dp := make([]int, amount+1)
    
    for i:= 0; i<len(dp); i++ {
        dp[i] = amount+1
    }
    
    dp[0] = 0
    
    for i:=1; i<=amount; i++ {
        for _,coin := range coins {
            if i - coin >= 0 {
                dp[i] = min(dp[i], 1 + dp[i-coin])
            }
        }
    } 
    
    if  dp[amount] == amount +1 {
       return -1
    }
    
    return dp[amount]
}

func min(a,b int) int {
    if a > b {
        return b
    }
    return a
}
