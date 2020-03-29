package main

func reverseString(s []byte)  {
	size := len(s)-1
	for i:=0;i<len(s)/2;i++{
		s[i],s[size-i] = s[size-i],s[i]
	}

}