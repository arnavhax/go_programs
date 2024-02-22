package main

import(
	"fmt"
)

func fibo(n int)int{
	if(n==1){
		return 0
	}
	if(n==2){
		return 1
	}
	return fibo(n-1)+fibo(n-2)
}
func main(){
	fmt.Println("Enter n for fibo")
	var n int
	fmt.Scan(&n)
	fmt.Printf("Fibo of %d is %d\n", n, fibo(n))

}