package main

//https://leetcode-cn.com/problems/valid-parentheses/

func isValid(s string) bool {
	sli := s[:]

	var stack = make([]byte,0)

	var m map[byte]byte
	m = make(map[byte]byte)
	m[')'] = '('
	m['}'] = '{'
	m[']'] = '['

	var line = []byte{')','}',']'}

	for i := 0; i < len(sli); i++{

		if inExist(sli[i],line) && len(stack) > 0{
			if m[sli[i]] == stack[len(stack) - 1]{
				stack = stack[:len(stack) - 1]
			}else {
				stack = append(stack, sli[i])
			}
		}else {
			stack = append(stack, sli[i])
		}
	}

	if len(stack) > 0{
		return false
	}
	return true
}

func inExist(tag byte, line []byte) bool {
	for _,v := range line{
		if v == tag{
			return true
		}
	}
	return false
}

func main()  {
	isValid("()")
}