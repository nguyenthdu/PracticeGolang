package main

import "fmt"

/* basic data structure upon which we'll define methods */
type employee struct {
	salary float32
}

/*
a method which will add a specified percent to an

	employees salary
*/
func (this *employee) giveRaise(pct float32) {
	this.salary += this.salary * pct
}

func main() {
	/* create an employee instance */
	var e = new(employee)
	e.salary = 100000
	/* call our method */
	e.giveRaise(0.04)
	fmt.Printf("Employee now makes %f", e.salary)
}
