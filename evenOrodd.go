package main

import(
	"fmt"
)

func main(){
	var inp int

	fmt.Println("Enter the number")
	fmt.Scan(&inp)

	if(inp%2!=0){
		fmt.Printf("%d is odd\n",inp)
	}else{
		fmt.Printf("%d is even\n",inp)
	}

}