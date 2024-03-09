package main

import(
	"fmt"
	"math"
)

// The above code defines an interface named shape with a method area that returns a float64 value.
// @property {float64} area - The `area` method is a method signature defined in the `shape` interface.
// It specifies that any type that implements the `shape` interface must also implement an `area`
// method that returns a `float64` value.
type shape interface{
	area() float64
}

type circle struct {
	x,y,radius float64
}

type rectangle struct {
	width,height float64
}

// The function `func(circ circle) area() float64` is defining a method named `area` for the `circle`
// struct. This method calculates and returns the area of a circle by using the formula for the area of
// a circle, which is πr^2 (Pi times radius squared). In this case, it calculates the area of the
// circle based on the radius value stored in the `circle` struct.
func(circ circle) area() float64 {
	return math.Pi * circ.radius * circ.radius

}
// The function `func(rect rectangle) area() float64` is defining a method named `area` for the
// `rectangle` struct. This method calculates and returns the area of a rectangle by multiplying its
// width and height.
func(rect rectangle) area() float64 {
	return rect.width * rect.height
}

// The function `getArea` takes a shape interface and returns the area of the shape by calling the
// `area` method on the shape.
func getArea(shape shape) float64{
	return shape.area()
}

func main() {
	c := circle{ x: 0, y:-8, radius: 5 }
	r := rectangle{ width:10,height: 10}

	fmt.Printf("Circle area %f\n",getArea(c))
	fmt.Printf("rectangle are %f\n",getArea(r))
}