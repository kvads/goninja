package main

import "fmt"

const msg string = `Hello World!`

func main() {
	fmt.Println(msg)
	var smsg string = `Hello again!`
	smsg = `Rewrite Hello again!`
	fmt.Println(smsg)
	tmsg := `Hello again! Again!`
	fmt.Println(tmsg)
}
