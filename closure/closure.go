// A function that remembers variables from its outer scope, even after that outer function has finished.
// Each closure has its own memory
package main

import "fmt"

func closure() func() int{
	cnt := 0
	return func() int{
		cnt += 1
		return cnt
	 }
}

func main(){
	count := closure()
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())

}