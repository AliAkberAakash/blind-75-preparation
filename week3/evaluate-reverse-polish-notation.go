func evalRPN(tokens []string) int {
    
    var stack []int
    
    for _,token := range tokens {
        if tokenIsOperand(token) {
            l := len(stack)
            num1 := stack[l-2]
            num2 := stack[l-1]
            stack = stack[:l-2]
            stack = append(stack,calculate(num1,num2,token))
        }else {
            num,_ := strconv.Atoi(token)
            stack = append(stack,num)
        }
    }
    
    return stack[0]
}

func tokenIsOperand(token string) bool {
    switch token {
        case "+":
            fallthrough
        case "-":
            fallthrough
        case "/":
            fallthrough
        case "*":
            return true
        default:
            return false
    }
    
    return false
}

func calculate(num1,num2 int, token string) int{
    switch token {
        case "+":
            return num1+num2
        case "-":
            return num1-num2
        case "/":
            return num1/num2
        case "*":
            return num1*num2
    }
    
    return 0
}
