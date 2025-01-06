package main

import "fmt"

func main() {
	printMessage(sayHello(`Кирилл`))
}

func printMessage(msg string) {
	fmt.Println(msg)
}

func sayHello(name string) string {
	return `Hello,` + name
}
