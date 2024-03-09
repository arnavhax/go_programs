// programs to display maps

package main

import (
	"fmt"
)

func main() {
	places:= map[string]string{"Nepal":"Kathmandu", "US":"Washington DC"}

	for keys := range places {
		fmt.Println(keys)
	}
}