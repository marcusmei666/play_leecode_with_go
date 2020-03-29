package main

import "fmt"

//https://leetcode-cn.com/problems/add-binary/

func addBinary(a string, b string) string {
	ans := ""
	ca := 0
	lena := len(a) -1
	lenb := len(b) -1

	for lena >= 0 || lenb >= 0{
		sum := ca//每次进入循环会记录上一次的进位

		if lena >= 0{
			sum += int(a[lena]-'0')
		}//如果长度 >=0 加上这个值
		if lenb >= 0{
			sum += int(b[lenb]-'0')
		}// 和 余 2 的结果为本位的值
		if sum%2 == 1{
			ans = "1" + ans
		}else {
			ans = "0" + ans
		}
		//进位的值为结果除2
		ca = sum/2
		lena--
		lenb--
	}
	if ca == 1 {
		ans = "1" + ans
	}
	return ans
}

func main(){
	//a := "111"
	b := '1'
	fmt.Println(int(b))
	//c := []byte(a)
	 //d := string(c)

	//fmt.Println(addBinary(a,b))
}