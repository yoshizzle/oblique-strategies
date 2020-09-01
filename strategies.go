package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    strategies := []string {
        "strategy 1",
        "stragegy 2",
        "strategy 3",
        "strategy 4",
    }

    fmt.Println(strategies[rand.Intn(len(strategies))])
}
