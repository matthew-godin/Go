package main
import "fmt"

func main() {
	//message := "Hello, world"
	/*var i int = 32
	var f float32 = float32(i)*/
	const (
		a = iota
		b
		c
	)
	fmt.Println(a + b + c)
}