package main
import ("fmt";"math")

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}
func area(r Rectangle) float64 {
	return r.width * r.height
}
func area_C(r Circle) float64 {
	return r.radius * r.radius * math.Pi
}
func main(){
	r1 := Rectangle{12, 2}
	c1 := Circle{10}
	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of c1 is: ", area_C(c1))
}
