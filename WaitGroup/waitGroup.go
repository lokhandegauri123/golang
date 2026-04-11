/*
	WaitGroup is used to wait for multiple goroutines to finish execution.

It comes from the package:
👉 sync

Imagine:

You sent 5 workers to do tasks 👷‍♂️
You don’t want to leave until all finish

👉 So you:

Count workers
Wait for all to report back
*/
package main

import (
	"fmt"
	"sync"
)

func printEle(i int, w *sync.WaitGroup){
	defer w.Done()
	
	fmt.Println(i)
	
}

func main(){
	var wg sync.WaitGroup
	// arr := []int{1,2,3}
	for i:=0 ; i<=10 ;i++{
		wg.Add(1)
		go printEle(i , &wg)
	}
	// printEle(arr, &wg)

	wg.Wait()
}