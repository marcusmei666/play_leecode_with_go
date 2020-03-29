package main

import "fmt"
func plusOne(digits []int) []int {
	var flag bool
	for i := len(digits) - 1; i >= 0; {
		if digits[i] + 1 > 9{
			flag = true
			digits[i] = 0
			i--

		}else {
				fmt.Println(digits)
				flag = false
				digits[i] = digits[i] + 1
				break
		}
	}
	if !flag {
		return digits
	}else {
		return append([]int{1}, digits...)
	}
}

func main()  {
	var a = []int{0}
	fmt.Println(plusOne(a))


}