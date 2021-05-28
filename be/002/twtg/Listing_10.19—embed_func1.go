package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func main() {
	c := new(Customer)
	c.Name = "Michael"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	// shorter:
	c1 := &Customer{"Michael2", &Log{"1 - Yes we can!"}}
	fmt.Println(c)
	fmt.Println(c1)

	c.Log().Add("2 - Yes we can + 1")

	fmt.Println(c.log)
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}
