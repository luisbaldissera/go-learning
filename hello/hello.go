package main

import (
    "fmt"
    "luisbaldissera/go-learning/greetings"
    "log"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)
    // A slice of names
    names := []string{"Lara", "Luis"}
    // Get a greeting message and print it.
    messages, err := greetings.Hellos(names)
    if (err != nil) {
        log.Fatal(err)
    }
    fmt.Println(messages)
}
