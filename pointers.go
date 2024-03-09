// stored in hexa decimal format
// &a gives location of a in memory

// we can name pointer in go and then refer them

package main
import (
	"fmt"
)
func swap(a,b int)(int,int){
	return b,a

}
func main() {
	// var intdata=20
	// var intPointer *int=&intdata

	// fmt.Println("What intData Stores",intdata)
	// fmt.Println("Address of intData",&intdata)
	// fmt.Println("What intPointer stores",intPointer)

	// *intPointer=30

	// fmt.Println("what intdata now stores",intdata)


	// swap without pointer in go
	// var a int=10
	// var b int=20
	// fmt.Println("a",a,"b",b)
	// a,b=swap(a,b)
	// fmt.Println("a",a,"b",b)
	// a,b=b,a
	// fmt.Println("a",a,"b",b)


}