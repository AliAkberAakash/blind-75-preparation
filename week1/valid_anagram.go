func isAnagram(s string, t string) bool {
    
    lookupTable := make(map[int]int)
    
    for _,c := range s {
        lookupTable[int(c)]++
    }
    
    for _,c := range t {
        lookupTable[int(c)]--
    }
    
    for _,val := range lookupTable {
        if val != 0 {
            return false
        }
    }
    
    return true
}
