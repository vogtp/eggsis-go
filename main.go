package main

import (
	"fmt"

	"github.com/vogtp/eggsis-go/pkg/cfg"
	"github.com/vogtp/eggsis-go/pkg/loop"
)


func main() {
	fmt.Printf("Starting %s...\n",cfg.AppName)
	loop.Start()
    // b:=4050.0/55
	// a:=3
	// name:="Lauri"
	// fmt.Printf("%v: a=%v b=%v\n",name,a,b)
	// for i:=a; a<10; i = i+a{
	// 	a=i+2
	// 	fmt.Printf("i=%v a=%v -> a+i=%v\n",i,a, add(a,i))
	// }
}

func add(x int, y int) int {
	x=x+3*y
	if x >20 {
		return -2
	}
	return x+y
}