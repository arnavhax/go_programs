package main
import(
	"fmt"
)

// type Vertex struct {
// 	Lat,Long float64
// }

// var m = map [string]Vertex{
// 	"Bell_Labs":Vertex{
// 		40.212,32.221,},
// 	"Google":Vertex{
// 		37.21212,-122.322,
// 	},
// }
func main(){
	// fmt.Println(m)
	// m["Nagpur"]=Vertex{22.212,21.211}
	// fmt.Println(m)


	// make function for memory allocation and and two value assignment in map

	m:=make(map [string]int)
	m["Answer"]=42
	v,err := m["Answer"]
	fmt.Println("Key value of answer",v,"Is answer key exist",err)
	delete(m,"Answer")
	v_new,err_new := m["Answer"]
	fmt.Println("Key value of answer",v_new,"Is answer key exist",err_new)
}
