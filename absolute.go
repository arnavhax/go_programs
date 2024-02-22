package main

import(
	"fmt"
)

func main() {
	var inp int
	fmt.Println("Enter a number: ")
	fmt.Scan(&inp)
	if(inp>=0){
		fmt.Printf("%d is in absolute form\n",inp)
	}else{
		fmt.Printf("%d is in not in absolute for and absolute form is %d\n",inp,-inp)

	}
}