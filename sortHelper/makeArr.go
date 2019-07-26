package sortHelper

import (
	"math/rand"
	"time"
)

func MakeIntArr (arr *[10]int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0 ;i < 10 ; i++{
		temp := rand.Intn(100)
		(*arr)[i] = temp
	}
	return
}

func MakeIntSli() (arr []int){
	arr = make([]int,10)
	rand.Seed(time.Now().UnixNano())
	for i := 0 ;i < 10 ; i++{
		temp := rand.Intn(100)
		arr[i] = temp
	}
	return
}

func MakeIntSliWithLen(len int)  (arr []int){
	arr = make([]int,len)
	rand.Seed(time.Now().UnixNano())
	for i := 0 ;i < len ; i++{
		temp := rand.Intn(100)
		arr[i] = temp
	}

	return
}
