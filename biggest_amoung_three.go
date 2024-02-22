package main

import(
	"fmt"
)

func main(){
	var num1,num2,num3 int=1,1,3

	if num1>num2 {
		if num1 >num3 {
			fmt.Printf("Largest of the three is %d\n",num1)
		} else{ // else must be on the same line as the end of if 
			fmt.Printf("Largest of the three is %d\n",num2)
		}
		
	} else {
		if num2>num3 {
			fmt.Printf("Largest of the three is %d\n",num3)
		}
	}
	//Not allowed to be on the same line as the end of if or else
	// if condition {

	// }
	// else {

	// }

	// if{
		// 	statement;
	// } else 
	//{


	// scanln takes only one value in a line if declared separately
	//scanln(&a, &b, &c, &d) will take four values in a single line

}