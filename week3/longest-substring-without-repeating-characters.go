func lengthOfLongestSubstring(s string) int {
    
    lookupTable := make(map[byte]bool)
    
    length := len(s)
    
    if length == 1 {
        return length
    }
    
    p1,p2 := 0,0
    maxTillNow := 0
    
    for p1<length && p2<length {
        _,ok := lookupTable[s[p2]]
        if ok {
            maxTillNow = max(maxTillNow,len(lookupTable))        
            p1++
            p2=p1
            lookupTable = make(map[byte]bool)
        }else {
            lookupTable[s[p2]] = true
            p2++
        }
        
    }
    maxTillNow = max(maxTillNow,len(lookupTable)) 
    return maxTillNow
}

func max(x,y int) int {
    if x>y {
        return x
    }
    return y
}
