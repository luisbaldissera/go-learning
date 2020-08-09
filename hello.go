package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }

}

func sayHello() {
    fmt.Println("Hello to all the readers that are present here in this tiny very long string text. Lorem ipsu det sit amet consectour ...")
}

func main() {
    go sayHello()
    go sayHello()
    go sayHello()
    go sayHello()
    go sayHello()
    sayHello()
    var i int
    fmt.Scanf("%d", &i)
    fmt.Println(i)
}
