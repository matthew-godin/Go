package main

/*import (
	"fmt"
	"bufio"
	"os"
	"strings"
)*/
/*import (
	"os"
	"net/http"
	"io"
)*/
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
	/*a := 42
	b := &a
	a = 27*/
	//c := new(int)
	//fmt.Println(*c)
	/*in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s + "!")*/
	/*http.HandleFunc("/", Handler)
	http.ListenAndServe(":3000", nil)*/
	arr := [3]string{"foo", "bar", "baz"}
	arr2 := arr
	arr[0] = "quux"
	fmt.Println(arr2)
}

/*func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)
}*/
