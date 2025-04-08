package vehicles

import (
	Layout "github.com/alphastigma101/Coconuts-At-Wars/layout"
)

type movement interface {
	Up()
	Down()
	Left()
	Right()
}

type plane struct {
	Health   int
	Armror   int
	Movement *movement
}
type Plane plane

type boat struct {
	Health   int
	Armror   int
	Movement *movement
}
type Boat boat

type car struct {
	Health   int
	Armror   int
	Movement *movement
}
type Car car

type truck struct {
	Health   int
	Armror   int
	Movement *movement
}
type Truck truck

func (p *Plane) Up() {
	zoom := Layout.GetTime()
	print(zoom)
}

func (B *Boat) Up() {

}

func (C *Car) Up() {

}

func (T *Truck) Up() {

}

func (p *Plane) Down() {

}

func (B *Boat) Down() {

}

func (C *Car) Down() {

}

func (T *Truck) Down() {

}

func (p *Plane) Left() {

}

func (B *Boat) Left() {

}

func (C *Car) Left() {

}

func (T *Truck) Left() {

}

func (p *Plane) Right() {

}

func (B *Boat) Right() {

}

func (C *Car) Right() {

}

func (T *Truck) Right() {

}
