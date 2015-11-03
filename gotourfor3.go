package main

import "fmt"

func main(){
	sum := 1
	for sum < 2000  {
		sum += sum
	}
	fmt.Println(sum)
}
