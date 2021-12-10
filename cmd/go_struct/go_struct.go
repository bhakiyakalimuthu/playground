package main

import "fmt"

type Cat struct {
	color string
}

type Dog struct {
	color string
}

func (d *Dog) changeColor() {
	d.color = "yellow"
}

func (c Cat) changeColor() {
	c.color = "pink"
}

func main() {
	funCat := Cat{"black"}
	funCat.changeColor()
	fmt.Println(funCat.color)

	funDog := Dog{"blue"}
	funDog.changeColor()
	fmt.Println(funDog.color)

	rose := Person{"rose"}
	s := SoccerPlayer{rose, "win a world cup"}
	fmt.Println(s.Talk())
	fmt.Println(s.Inspire())
	fmt.Println(s.lifeGoal)
}

type Person struct {
	name string
}

func (p Person) Talk() string {
	return "My name is " + p.name
}

type SoccerPlayer struct {
	Person
	lifeGoal string
}

func (s SoccerPlayer) Inspire() string {
	s.lifeGoal = "win a olympic"
	return "My life goal is to " + s.lifeGoal
}
