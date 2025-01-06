package main

import "fmt"

func main() {
	printMessage(`Текст 1`)
	printMessage(`Текст 2`)
	printMessage(`Текст 3`)
}

func printMessage(msg string) {
	fmt.Println(msg)
}
