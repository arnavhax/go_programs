package main
import "fmt"

func main() {
	primeNumbers:=[] int {2,3,5,7}

	numbers:= [] int{1,2,3}

	copy(numbers,primeNumbers)

	fmt.Print(numbers)
}