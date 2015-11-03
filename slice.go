package main
import "fmt"
func main() {
	var arr [10]int
	arr[0]=42
	arr[1]=49
	fmt.Printf("The first element is %d\n", arr[0])
	fmt.Printf("The second element is %d\n", arr[1])
	a := [3]int{1,2,3}
	b := [10]int{1,2,3}
	c := [...]int{4,5,6}
	fmt.Printf("The second element of a is %d\n", a[1])
	fmt.Printf("The second element of b is %d\n", b[1])
	fmt.Printf("The second element of c is %d\n", c[1])
	var array = [10]byte{'a','b','c','d','e','f','g','h','i','j'}
	var aSlice,bSlice []byte
	aSlice = array[:3]

	fmt.Println(aSlice)
	bSlice = array[3:5]
	fmt.Println(bSlice)
}
