package tfmt

import "fmt"
import "io"

func PrintV() {
	s := &Student{"hbdrawn", 26, "男"}
	fmt.Printf("%+v\n", s)
	fmt.Printf("%#v\n", s)
	fmt.Printf("%v\n", s.age)
	fmt.Printf("%T\n", s)
	fmt.Printf("%%\n", s)

}

func PrintInt() {
	s := new(Student)
	s.age = 55
	s.name = "father	"
	s.sex = "男"
	fmt.Printf("%b\n", s.age)
	fmt.Printf("%d\n", s.age)
	fmt.Printf("%o\n", s.age)
	fmt.Printf("%c\n", s.age)
	fmt.Printf("%q\n", s.age)
	fmt.Printf("%x\n", s.age)
	fmt.Printf("%X\n", s.age)
	fmt.Printf("%U\n", s.age)
}

func PrintFloat() {
	s := &Student{"hbdrawn", 57.90, "man"}
	fmt.Printf("%b\n", s.age)
	fmt.Printf("%e\n", s.age)
	fmt.Printf("%E\n", s.age)
	fmt.Printf("%g\n", s.age)
	fmt.Printf("%G\n", s.age)

}

func PrintSTR() {
	s := &Student{"hbdrawn", 57.90, "man"}
	fmt.Printf("%s\n", s.name)
	fmt.Printf("%+q\n", s.name)
	fmt.Printf("% x\n", s.name)
	fmt.Printf("% X\n", s.name)

}

func Scan() {
	var tmp string
	var in string = "this is a test"
	num, err := fmt.Scanln(&in)
	if err != nil {
		if err != io.EOF {
			fmt.Println(err)
			return
		}
	}
	fmt.Printf("num:%d  tmp:%q\n", num, tmp)
}

type Student struct {
	name string
	age  float64
	sex  string
}
