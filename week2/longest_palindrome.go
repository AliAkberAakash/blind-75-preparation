func longestPalindrome(s string) int {
    
    lookupTable := make(map[rune]int)
    
    sum := 0
    odd := 0
    
    for _,r := range s {
        lookupTable[r]++
        if lookupTable[r]%2==0 {
            sum+=2
        }
    }
    
    for _,r := range s {
        if lookupTable[r]%2==1 {
            odd = 1
            break
        }
    }
    
    
    return sum+odd
}
