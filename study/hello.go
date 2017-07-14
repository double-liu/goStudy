// go "hello world" 示例
package main

import (
	"fmt"
)

var name = "hello word"
var price = 333.222
var age int64 = 33
var a [4]int = [4]int{2, 3, 4, 5}

type   int1  int
type person struct {
	name string
	age  int
	int1
	int
}

var  p  =  person{name:"aa"}
var zs  person
type nameBlock func(str string) string

func main() {
	s := []string{"aa", "bb"}
	fmt.Printf("slice%d\n", s[0])
	switch name {
	case "aa":
		fmt.Printf("name is aa")
		fallthrough
	case "hello word":
		fmt.Printf("name is hello word")
	}
	ll, _ := otherName()
	fmt.Printf(ll)
	var more = moreNmae("11", "22", "33")
	fmt.Printf(more[2])

	block(func(str string) string {
		return str + "AA"
	})
	zs.name = "张三"
	zs.age  = 33
	zs.int  = 33
	zs.int1 = 44
}

func otherName() (name string, err error) {
	return "张三", nil
}

func moreNmae(arg ...string) []string {
	return arg
}

func block(block nameBlock) {
	name := block("block")
	fmt.Printf(name)
}



