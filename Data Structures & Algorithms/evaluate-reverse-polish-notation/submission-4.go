func evalRPN(tokens []string) int {
    stack := make([]int, 0)
    for _, op := range tokens {
        if op == "+" || op == "-" || op == "*" || op == "/" {
            h := len(stack) - 1
            calc := 0
            if op == "+" {
                calc = stack[h-1] + stack[h]
            } else if op == "-" {
                calc = stack[h-1] - stack[h]
            } else if op == "*" {
                calc = stack[h-1] * stack[h]
            } else {
                calc = stack[h-1] / stack[h]
            }
            stack = stack[:h-1]
            stack = append(stack, calc)
        } else {
            num, _ := strconv.Atoi(op)
            stack = append(stack, num)
        }
    }
    return stack[len(stack)-1]
}