package main

import "fmt"

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
	/*arr := [3]string{"foo", "bar", "baz"}
	arr2 := arr
	arr[0] = "quux"
	fmt.Println(arr == arr2)*/
	/*var s = []int{1, 2, 3}
	fmt.Println(s[1])
	s = append(s, 5, 10, 15)
	fmt.Println(s)
	s = slices.Delete(s, 1, 3)
	fmt.Println(s)*/
	/*m := map[string]int{"foo": 1, "bar": 2}
	m["bar"] = 99
	delete(m, "foo")
	fmt.Println(m)
	v, ok := m["foo"]
	fmt.Println(v, ok)*/
	/*var s struct {
		name string
		id   int
	}
	s.name = "Arthur"
	fmt.Println(s)*/
	/*type myStruct struct {
		name string
		id   int
	}
	s := myStruct{
		name: "Arthur",
		id:   42}
	s2 := s
	s.name = "Tricia"
	fmt.Println(s == s2)*/
	/*i := 1
	for {
		fmt.Println(i)
		i += 1
		break
	}
	i = 1
	for i < 3 {
		fmt.Println(i)
		i += 1
	}
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}*/
	/*m := map[string]int{"foo": 1, "bar": 2}
	for _, value := range m {
		fmt.Println(value)
	}*/
	/*if i := 5; i < 5 {
		fmt.Println("Less than 5")
	} else if i < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("At least 10")
	}*/
	/*switch i := 7; i {
	case 1:
		fmt.Println("first case")
	case 2 + 3, 2*2 + 3:
		fmt.Println("second case")
	default:
		fmt.Println("default case")
	}*/
	switch i := 4; {
	case i < 5:
		fmt.Println("less than 5")
	case i < 10:
		fmt.Println("less than 10")
	default:
		fmt.Println("at least 10")
	}
}

/*func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)
}*/
