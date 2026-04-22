package main

import "fmt"

const englishGreetingPrefix = "Hello, "

func Hello(name string) string {
	return englishGreetingPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}