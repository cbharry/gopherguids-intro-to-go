package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	var first int8 = 1
	var second int = 2
	fmt.Printf("%d + %d = %d\n", first, second, int(first)+second)
	d := []string{"a", "b", "c"}
	fmt.Printf("%#v\n", d)
	goString := "Go"
	fmt.Printf("Let's %[1]v, %[1]v, %[1]v!\n", goString)
	pi := 3.14159265359
	fmt.Printf("PI %.2f\n", pi)
}
