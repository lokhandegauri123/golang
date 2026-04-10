

package main
import "fmt"

func pointer(num *int) int{
	*num = 5
	return *num	
}

func main(){
	var num int = 1
	fmt.Println("num before change",num)
	fmt.Println(pointer(&num))
	fmt.Println("num before change",num)


}