package domain

import (
	"mars-rover/src/exception"
	"mars-rover/src/valueObject"
)

type Plateau struct {
	Area   valueObject.Point
	Rovers []*Rover
}

func NewPlateau(area valueObject.Point) Plateau {
	return Plateau{Area: area}
}

func (p *Plateau) AddRover(rover *Rover) error {
	for _, r :=range p.Rovers {
		if r.Position == rover.Position {
			return exception.AlreadyOccupied
		}
	}
	p.Rovers = append(p.Rovers, rover)
	return nil
}
