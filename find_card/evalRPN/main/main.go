package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack = make([]int,0)
	var res int
	for i := 0; i < len(tokens); i++{
		if isNum([]byte(tokens[i])){
			tempint, err := strconv.Atoi(tokens[i])
			if err != nil{
				panic(err)
			}
			stack = append(stack,tempint)
		}else {
			//取出两个数字
			if len(stack) < 2{
				panic("stack is nil")
			}

			res = 0
			switch tokens[i] {
			case "+":
				res = stack[len(stack)-2] + stack[len(stack)-1]
			case "-":
				res = stack[len(stack)-2] - stack[len(stack)-1]
			case "*":
				res = stack[len(stack)-2] * stack[len(stack)-1]
			case "/":
				res = stack[len(stack)-2] / stack[len(stack)-1]
			}
			stack = stack[:len(stack)-2]
			stack = append(stack,res)
		}

	}

	return int(stack[0])
}

func isNum(temp []byte) bool {
	if temp[len(temp)-1] <='9' && temp[len(temp)-1] >= '0'{
		return true
	}
	return false
}

func main()  {
	var s = []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}
	fmt.Println(evalRPN(s))
	//var stack = make([]int,0)
	//s := "-11"
	//tempint, err := strconv.Atoi(s)
	//if err != nil{
	//	panic(err)
	//}
	//
	//stack = append(stack,tempint)
	//stack = append(stack,tempint)
	//fmt.Println(tempint,stack)
	//
	//res := stack[0] + stack[1]
	//
	//fmt.Println(res)
	//
	//var a int
	//
	//a = -11
	//fmt.Println(a)
}