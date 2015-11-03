package main
import "fmt"

func add(x, y int) int {
	return x + y
}

func main(){
	fmt.Println(add(42,13))
	fmt.Printf("The sum is :%d\n",add(42,13))
}
