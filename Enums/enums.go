package main
import "fmt"

type OrderStatus string

const(
	// iot used only for integer
	/* 
	Received OrderStatus= iota
	Confirmed 
	Prepared 
	Delivered
	 */

	//  
	Received OrderStatus= "received"
	Confirmed ="confirmed"
	Prepared = "prepared"
	Delivered = "delivered"
	
	
)

func changeStatus(status string){
	fmt.Println(status)
}
func main(){
	changeStatus(Prepared)
}