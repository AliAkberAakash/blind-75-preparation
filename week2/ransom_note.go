func canConstruct(ransomNote string, magazine string) bool {
    lookupTable := make(map[rune]int)
    
    for _,c := range ransomNote {
        lookupTable[c]++
    }
    
    for _,c := range magazine {
        if _, ok := lookupTable[c]; ok {
            lookupTable[c]--
        }
    }
    
    for _,val := range lookupTable {
        if val > 0 {
            return false
        }
    }
    
    return true
}
