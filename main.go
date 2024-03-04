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
/*import (
	"fmt"
	"program/menu"
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
	/*switch i := 4; {
	case i < 5:
		fmt.Println("less than 5")
	case i < 10:
		fmt.Println("less than 10")
	default:
		fmt.Println("at least 10")
	}*/
	/*fmt.Println("main 1")
	defer fmt.Println("defer 1")
	fmt.Println("main 2")
	defer fmt.Println("defer 2")*/
	/*fmt.Println("main 1")
	func1()
	fmt.Println("main 2")*/
	/*for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				fmt.Println(i, j)
				if i+j == 15 {
					goto myLabel
				}
			}
		}
	myLabel:
		fmt.Println("End of program")*/
	//greet("a", "b", "c")
	/*name, otherName := "Name", "Other name"
	fmt.Println(name)
	fmt.Println(otherName)
	myFunc(name, &otherName)
	fmt.Println(name)
	fmt.Println(otherName)*/
	//fmt.Println(divide(45, 0))
	//fmt.Println(menu.Menu)
	fmt.Println(myInt.isEven(3))
}

type myInt int

func (i myInt) isEven() bool {
	return int(i)%2 == 0
}

/*func greet(name1 string, name2 string) {
	fmt.Println(name1)
	fmt.Println(name2)
}*/

/*func greet(name1, name2 string) {
	fmt.Println(name1)
	fmt.Println(name2)
}*/

/*func greet(names ...string) {
	for _, n := range names {
		fmt.Println(n)
	}
}*/

/*func myFunc(name string, otherName *string) {
	name = "New name"
	*otherName = "Other new name"
}*/

/*func divide(l, r int) (int, bool) {
	if r == 0 {
		return 0, false
	}
	return l / r, true
}*/

/*func divide(l, r int) (result int, ok bool) {
	if r == 0 {
		return
	}
	result = l / r
	ok = true
	return
}*/

/*func func1() {
	defer func() {
		fmt.Println(recover())
	}()
	fmt.Println("func 1 1")
	panic("uh-oh")
	fmt.Println("func 1 2")
}*/

/*func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)
}*/
