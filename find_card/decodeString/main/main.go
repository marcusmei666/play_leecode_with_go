package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/decode-string/

func decodeString(s string) string {
	sli := []byte(s)
	intstack := make([]int,0)
	strtack := make([]string,0)
	var re string
	for i := 0; i < len(sli); i++{
		if sli[i] >= '0'  && sli[i] <= '9'{
			intstack = append(intstack, int(sli[i])-48)
		}else if sli[i] == ']'{
			var temp,res string
			for{
				if strtack[len(strtack)-1:][0] == "["{
					strtack = strtack[:len(strtack)-1]
					break
				}
				temp = strtack[len(strtack)-1:][0] + temp
				strtack = strtack[:len(strtack)-1]
			}

			for j := 0; j < intstack[len(intstack)-1:][0];j++{
				res = res+temp
			}
			intstack = intstack[:len(intstack)-1]
			strtack = append(strtack,res)

		}else {
			strtack = append(strtack,string(sli[i]))

		}
	}

	for i := 0; i < len(strtack); i++{
		re += strtack[i]
	}
	return re
}

func main()  {
	fmt.Println(decodeString("100[leetcode]"))
}