package main

import (
	"bom-studies-go/greetings"
	"bom-studies-go/stringutil"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(greetings.Hello("bom"))
	fmt.Println(stringutil.ToUpper("bom to upper"))
}
