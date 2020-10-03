package main

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	ss := []byte(s)
	stack := make([]byte, 0, len(s))
	for _, v := range ss {
		switch v {
		case '[':
			fallthrough
		case '{':
			fallthrough
		case '(':
			stack = append(stack, v)
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} //swtich
	} //for
	return len(stack) == 0
}
