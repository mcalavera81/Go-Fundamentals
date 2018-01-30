package main

import "fmt"

func main() {
	var command string
	var valid bool

	fmt.Printf("command %T %q\n",command,command)
	fmt.Printf("valid %T %v\n",valid,valid)

	foo := 5
	bar := true

	fmt.Printf("foo %T %v\n",foo,foo)
	fmt.Printf("bar %T %v\n",bar,bar)

}