package checkSlice

import (
	"fmt"
)

type per struct {
	name    string
	age, id int
}

/**
*	结构体基础
 */
var person = make(map[string]per)

// var person = make(map[string]*per)
var persons = []per{
	{name: "1111", age: 12, id: 1},
	{name: "Tom", age: 12, id: 1},
	{name: "Angel", age: 12, id: 1},
	{name: "Sipack", age: 14, id: 3},
	{name: "&&&&", age: 13, id: 2},
}

func GetStructList() {
	/**
	*	for循环结构体
	*	每次循环循环内容赋值变量i
	*	当使用指针时修改了原始值
	 */
	for _, i := range persons {
		person[i.name] = i
		// person[i.name] = &i
	}

	for p, r := range person {
		fmt.Println(p, "=>", r.name)
	}

	fmt.Println("\n", "GetStructList=============================")
	fmt.Println(person)
	fmt.Println(persons)
	fmt.Println("GetStructList=============================", "\n")
}

func CreateArr() {
	b := []int{1, 2, 3, 4, 5}

	arrString := [3]string{"one", "two"}
	AllArrString := arrString[:]
	item := arrString[2]
	first := arrString[:1]
	last := arrString[0:]

	a := [5]int{1, 2, 3, 4, 5}
	capArr := a[:2]
	fmt.Println(len(capArr), cap(capArr), "capArr")

	first[0] = "t5hree"

	fmt.Println("\n", "CreateArr=============================")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(capArr[0], capArr, len(capArr), cap(capArr), "capArr")
	fmt.Println(AllArrString)
	fmt.Println(item)
	fmt.Println(first)
	fmt.Println(last)
	fmt.Println("CreateArr=============================", "\n")
}
