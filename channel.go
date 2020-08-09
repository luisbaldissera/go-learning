package main

import (
    "fmt"
    "time"
)

func produce(c chan int) {
    i := 0
    for {
        c <- i
        i++
        time.Sleep(500 * time.Millisecond)
    }
}

func consume(c chan int) {
    for {
        fmt.Println("Consumed:", <-c)
    }
}

func main() {
    ch := make(chan int)
    go produce(ch)
    consume(ch)
}
