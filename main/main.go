package main
import  "fmt"

func fact_recursion(x int) (res int){
	if x > 0{
		res = x * fact_recursion(x -1)
	} else{
		res = 1
	}
	return
}

func main(){
	fmt.Println(fact_recursion(4))
}