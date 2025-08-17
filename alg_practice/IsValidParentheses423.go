package main

import "fmt"

/*
*给定一个字符串所表示的括号序列，包含以下字符： '('、')'、'{'、'}'、'[' 和 ']'， 判定是否是有效的括号序列，有效括号序列满足以下条件：

左括号必须用相同类型的右括号闭合
左括号必须以正确的顺序闭合
每个右括号都有一个对应的相同类型的左括号
*/
func IsValidParentheses(s string) bool {
	// write your code here
	l := len(s)
	if l == 0 {
		return true
	}
	var brackalMap = map[byte]byte{'}': '{', ']': '[', ')': '('}
	var stack []byte
	for i := 0; i < l; i += 1 {
		ch := s[i]
		if ch == '{' || ch == '[' || ch == '(' {
			stack = append(stack, ch)
			continue
		}
		lastIndex := len(stack) - 1
		if len(stack) == 0 || brackalMap[ch] != stack[lastIndex] {
			return false
		}
		stack = stack[:lastIndex]
	}
	return len(stack) == 0

}

func main() {
	fmt.Println(IsValidParentheses("()"))
}
