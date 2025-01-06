package main

import "fmt"

func main() {
	printMessage(sayHello(`Кирилл`, 25))
}

func printMessage(msg string) {
	fmt.Println(msg)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf(`Привет, %s! Тебе %d лет!`, name, age)
}
