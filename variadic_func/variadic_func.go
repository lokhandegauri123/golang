package main

import "fmt"
	

func sum(num ...int) int{
	total := 0
	for _, val := range num{
		total = total + val
	}

	return total
}

func main(){
	nums := []int{3,23,34,2}
	fmt.Println(sum(nums...))
}