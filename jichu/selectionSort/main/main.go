package main

import (
	"fmt"
	"play_leecode_withgo/sortHelper"
)

//选择排序

func selectionSort1()  {
	var sli = sortHelper.MakeIntSli()[:]
	temp := len(sli)
	var minIndex int
	for i:= 0; i< temp; i++ {
		//寻找[i，j）中最小值的下标
		minIndex = i
		for j:= i + 1; j< temp; j++ {
			if sli[j] < sli[minIndex]{
				minIndex = j
			}
		}
		sli[i],sli[minIndex] = sli[minIndex],sli[i]
	}
	fmt.Println(sli)
}

func main()  {
	selectionSort1()
}