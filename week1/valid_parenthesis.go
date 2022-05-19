func isValid(s string) bool {
    lookupTable := make(map[rune]rune)
    lookupTable[')']='('
    lookupTable['}']='{'
    lookupTable[']']='['
    
    var stack []rune
    
    for _,char := range s {
        l := len(stack)
        if _,ok := lookupTable[char]; ok  {
            if l>0 && stack[l-1] == lookupTable[char] {
                stack = stack[:l-1]
            }else{
                return false
            }
        }else{
            stack = append(stack,char)
        }
    }
    
    return len(stack) == 0
}
