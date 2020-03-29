package main

import "fmt"

func dailyTemperatures(T []int) []int {
	length := len(T);
	result := make([]int,length)

	//从右向左遍历
	for i := length - 2; i >= 0; i-- {
		// j+= result[j]是利用已经有的结果进行跳跃
		for j := i + 1; j < length; j+= result[j] {
			if T[j] > T[i] {
				result[i] = j - i;
				break;
			} else if result[j] == 0 { //遇到0表示后面不会有更大的值，那当然当前值就应该也为0
				result[i] = 0;
				break;
			}
		}
	}

	return result;
}

func main()  {
	var a = []int{73,74,75,71,69,72,76,73}
	fmt.Println(dailyTemperatures(a))
}