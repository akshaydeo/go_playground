package main 

import "fmt"

// method to initialize a sequence and increment
func initSeq() func() int {
	i := 0
	return func () int {
		i += 1
		return i
	}
}

// main function
func main() {
	intSeq := initSeq()
	fmt.Println(intSeq())
	fmt.Println(intSeq())
	fmt.Println(intSeq())
	fmt.Println(intSeq())
}
