package main

import(
	"fmt"
)

func main() {
	var fibo func(n int) int
	fibo=func(n int) int{
		if(n<2){
			return n
		}
		return fibo(n-1)+fibo(n-2)
	}
	fmt.Println(fibo(9)) // 34
}

//homework print sum of n terms usinf anonymous function