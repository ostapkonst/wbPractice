// 14. Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
package main

import (
	"fmt"
	"reflect"
)

func getType(v interface{}) string {
	// TypeOf returns the reflection [Type] that represents the dynamic type of i. If i is a nil interface value, TypeOf returns nil.
	// Kind returns the specific kind of this type.
	// String returns the name of k.
	return reflect.TypeOf(v).Kind().String()
}

func main() {
	fmt.Printf(
		"int: %t, string: %t, bool: %t, channel: %t\n",
		getType(1) == "int",
		getType("str") == "string",
		getType(true) == "bool",
		getType(make(chan int)) == "chan",
	)
}
