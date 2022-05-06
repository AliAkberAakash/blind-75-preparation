func addBinary(a string, b string) string {
    carry := '0'
    
    lenA := len(a)-1
    lenB := len(b)-1
    
    var sb strings.Builder
        sb.WriteString("")
    
    for lenA>=0 && lenB>=0 {
        c := add(rune(a[lenA]),rune(b[lenB]), &carry)
        sb.WriteRune(c)
        lenA--
        lenB--
    }
    
    for lenA >= 0 {
        c := add(rune(a[lenA]),'0', &carry)
        sb.WriteRune(c)
        lenA--
    }
    
    for lenB >= 0 {
        c := add('0',rune(b[lenB]), &carry)
        sb.WriteRune(c)
        lenB--
    }
    
    if carry == '1' {
         sb.WriteRune(carry)
    }
    
    return reverse(sb.String())
}

func add(a rune, b rune, carry *rune) rune{
    if a=='0' && b=='0'{
        if *carry == '0' {
            return '0'
        }else{
            *carry = '0'
            return '1'
        }
    }
    if (a=='0' && b=='1') || (b=='0' && a=='1'){
        if *carry == '0' {
            return '1'
        }else{
            *carry = '1'
            return '0'
        }
    }
    if a=='1' && b=='1'{
        if *carry == '0' {
            *carry = '1'
            return '0'
        }else{
            *carry = '1'
            return '1'
        }
    }
    
    return '0'
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
