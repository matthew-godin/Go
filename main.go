package main
import "fmt"

func main() {
	//message := "Hello, world"
	/*var i int = 32
	var f float32 = float32(i)*/
	/*const (
		a = iota
		b
		c
	)*/
	a := 42
	b := &a
	a = 27
	fmt.Println(*b)
}