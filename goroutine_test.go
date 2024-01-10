package go_routines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoRoutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("This is printed from the test")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoRoutine(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
