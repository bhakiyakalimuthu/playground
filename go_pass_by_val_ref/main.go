package main

import "fmt"

type person struct {
	firstName, secondName string
}

func changeSecondName(p person, name string) {
	p.secondName = name

}
func changeSecondNameRef(p *person, name string) {
	p.secondName = name

}

func newPerson() person {
	p := person{
		firstName:  "kavin",
		secondName: "mugil",
	}

	return p
}

func (p person) changeSecondName(name string) {
	p.secondName = name

}

func newPersonRef() *person {
	p := &person{
		firstName:  "kavin",
		secondName: "mugil",
	}

	return p
}

func (p *person) changeSecondNameRef(name string) {
	p.secondName = name

}
func main() {

	p1 := person{
		firstName:  "kavin",
		secondName: "mugil",
	}

	changeSecondName(p1, "kavin")

	fmt.Printf("\n p1:%v", p1)

	p2 := &person{
		firstName:  "kavin",
		secondName: "mugil",
	}
	changeSecondNameRef(p2, "kavin")
	fmt.Printf("\n p2:%v", p2)

	p3 := newPerson()

	p3.changeSecondName("chellam")
	fmt.Printf("\n p3:%v", p3)

	p4 := newPersonRef()

	p4.changeSecondNameRef("chellam")
	fmt.Printf("\n p4:%v", p4)
}
