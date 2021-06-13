package main

import (
	"fmt"
	"strconv"
)

type T struct {
	a int
	b float32
	c string
}

func main() {
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
}

func (t *T) String() string {
	b := fmt.Sprintf("%f", t.b)
	c := fmt.Sprintf("%#v", t.c)
	return strconv.Itoa(t.a) + " / " + b + " / " + c
}
