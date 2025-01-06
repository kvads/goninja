package main

import (
	"fmt"
	"reflect"
)

const msg string = `Hello World!`

func main() {
	fmt.Println(msg)
	var smsg string = `Hello again!`
	smsg = `Rewrite Hello again!`
	fmt.Println(smsg)
	tmsg := `Hello again! Again!`
	fmt.Println(tmsg)

	fmt.Println(reflect.TypeOf(msg))
}
