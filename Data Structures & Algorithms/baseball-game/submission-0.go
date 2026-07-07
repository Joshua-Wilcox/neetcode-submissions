func calPoints(operations []string) int {
	
	stack := make([]int, 0)

	for _, op := range operations {
		n := len(stack)
		fmt.Println(stack)
		if op == "+" {
			stack = append(stack, stack[n-2]+stack[n-1])
		} else if op == "D" {
			stack = append(stack, stack[n-1] * 2)
		} else if op == "C" {
			stack = stack[:len(stack)-1]
		} else {
			thisNum, _ := strconv.Atoi(op)
			stack = append(stack, thisNum)
		}
	}

	res := 0
	for _, num := range stack {
		res += num
	}
	return res

}
