package main

import (
	"fmt"
	"reflect"
	"runtime"
)

// Empty ...
type Empty struct{}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println("bs-gen running on " + filename)

	fmt.Println(reflect.TypeOf(Empty{}).PkgPath())

}
