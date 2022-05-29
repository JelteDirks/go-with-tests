package main

import (
	"fmt"
)

func main() {
	fmt.Println("main")
}

func Hello(name string) string {
    if name == "" {
        name = "world"
    }
	return "Hello, " + name + "!"
}
