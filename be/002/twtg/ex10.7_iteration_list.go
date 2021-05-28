package main

import "container/list"

//A method and the type on which it acts must be defined in the same package
/*
func (p *list.List) Iter() {

}
*/

type myList struct {
	list.List //anonymous field
}

func (p myList) Iter() {
	return
}

func main() {
	lst := new(myList)
	for _ = range lst.Iter() {

	}
}
