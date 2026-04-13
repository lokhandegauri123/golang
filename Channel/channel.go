/* package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "Message from channel 1"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "Message from channel 2"
    }()

    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    }
} */

// Buffered Channel
/* A buffered channel is a channel that can store multiple values before they are received.

It allows sending data without immediate receiver
Real-life Analogy

Think of a queue (line) 🧾:

People can stand in line (buffer)
You don’t need to serve immediately

Important Behavior ⚠️
✅ Send does NOT block (until buffer full) */

package main

import "fmt"

func main() {
    ch := make(chan int, 2)

    ch <- 10
    ch <- 20

    fmt.Println(<-ch)
    fmt.Println(<-ch)
}