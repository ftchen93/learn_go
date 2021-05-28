package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	Log  //embeded Log in Customer type
}

func main() {
	c := &Customer{"Michael2", Log{"1 - Yes we can!"}}
	fmt.Println(c)

	c.Add("2 - Yes we can + 1")

	fmt.Println(c)
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log)
}
