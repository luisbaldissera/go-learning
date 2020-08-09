package main

import (
    "fmt"
    "time"
)

func produceInt(c chan int) {
    for i := 0 ; i < 100 ; i++{
        c <- i
        val := 150 + i % 50
        time.Sleep(time.Duration(val) * time.Millisecond)
    }
    close(c)
}

func produceString(s chan string) {
    for i := 0; i < 20; i++ {
        s <- "new string"
        time.Sleep(500 * time.Millisecond)
    }
    close(s)
}

func consume(s chan string, c chan int) {
    var str string
    var x int
    var ok bool
    for {
        select {
        case str, ok = <-s:
            if (ok) {
                fmt.Println("String:", str)
            }
        case x, ok = <-c:
            if (ok) {
                fmt.Println("Int:", x)
            }
        default:
            fmt.Println(".")
            time.Sleep(50 * time.Millisecond)
        }
    }
}

func main() {
    ch := make(chan int, 10)
    st := make(chan string, 10)
    go produceInt(ch)
    go produceString(st)
    consume(st, ch)
}
