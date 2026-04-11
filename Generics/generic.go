/* 
Generics allow you to write functions or data structures that work with multiple types.

Instead of writing the same code for int, string, etc.,
you write it once and reuse it.
*/ 

package main
import "fmt"

// func showElement[T any](sli[] T) {
// 	for _,val := range sli{
// 		fmt.Println(val)
// 	}
// }

// or

func showElement[T any](sli []T) {
	for _,val := range sli{
		fmt.Println(val)
	}
}
type stack[T any] struct{
	element []T
}
func main(){
	// num := []int{1,2,3,4,5}
	// str := []string{"gauri","hii","hello"}
	myStack := stack[int]{
		element: []int{1,2,3,4,5},
	}
	showElement(myStack.element)
}  

