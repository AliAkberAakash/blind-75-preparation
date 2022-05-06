func majorityElement(nums []int) int {
    lookupTable := make(map[int]int)
    
    for _,num := range nums {
        lookupTable[num]++
    }
    
    n := len(nums)/2
    
    for key,val := range lookupTable {
        if val > n {
            return key
        }
    }
    
    return -1
}
