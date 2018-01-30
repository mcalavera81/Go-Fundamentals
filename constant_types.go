package main

import "fmt"

const (
	PI       = 3.14159
	week     = 7
	freezing = int64(32)
)

func main() {
	i := 10
	j := int32(20)
	fmt.Println(week + i)
	fmt.Println(week + j)

	// This will not work as it is already typed
	// fmt.Println(freezing + i)

	// This will not work either
	//fmt.Println(i + PI)

}
