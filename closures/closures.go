package main 

import "fmt"


type Int uint64

type complexNumber struct{
	rel Int
	img Int
}



type shape interface{
	area() Int
	volume() Int
}

type square struct{
	side Int
}

func (s square) area() Int {
	return s.side*s.side
}

func (s square) volume() Int {
	return s.side*s.side
}

func printArea(s shape){
	fmt.Println(s.area())
}

func (cnum *complexNumber) print(){
	fmt.Println(cnum.rel,"+",cnum.img,"i")
}

// method to initialize a sequence and increment
func initSeq() (uint8, func() (Int,Int)) {
	i := Int(2)
	return uint8(2), func () (Int, Int) {
		power := i
		i *= i
		return power,i
	}
}

// method that counts number of params and returns it
func countParams(nums ...int) int{
	return len(nums)
	
}


func channel1(){
	messages := make(chan complexNumber)

	go func() {messages <- complexNumber{25,26}}()
	go func() {messages <- complexNumber{28,29}}()

	msg := <-messages
	msg.print()
	msg1 := <-messages
	msg1.print()
}

func ping(channel chan<- string, message string){
	fmt.Println("in ping")
	channel <- message	
}

func pong(channel <-chan string){
	fmt.Println("in pong")
	msg := <-channel
	fmt.Println(msg)
}

// main function
func main() {
	channel := make(chan string,1)
	ping(channel, "message")	
	pong(channel)
}