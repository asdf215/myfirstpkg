package main

//"reflect"
import (
	"fmt"
)

func typeCheck(i interface{}) {
	fmt.Printf("%T \n", i)
}

func main() {
	typeCheck(1)
}
