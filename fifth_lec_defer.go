// use of defer keyword
package main

import(
	"fmt"
)

func main(){

	defer fmt.Println("Starting")
	defer fmt.Println("mid")
	fmt.Println("Last")
}