package domain

import (
	"fmt"
	"mars-rover/constants"
	"mars-rover/src/exception"
	"mars-rover/src/valueObject"
)

type Rover struct {
	Plateau   Plateau
	Position  valueObject.Position
	Direction valueObject.Direction
}

func NewRover(plateau Plateau, position valueObject.Position, direction valueObject.Direction) (*Rover, error) {

	if !position.IsInsideArea(plateau.Area) {
		return nil, exception.IsNotOnMars
	}

	return &Rover{
		Plateau:   plateau,
		Position:  position,
		Direction: direction,
	}, nil
}

func (r *Rover) Move(move valueObject.Move) (err error) {
	switch move {
	case constants.MoveForward:
		err = r.moveForward()
	case constants.MoveRight:
		err = r.turnRight()
	case constants.MoveLeft:
		err = r.turnLeft()
	}

	return err
}

func (r *Rover) moveForward() (err error) {
	var isInsideArea bool
	switch r.Direction {
	case constants.DirectionNorth:
		r.Position.AddY()
		isInsideArea = r.Position.IsInsideArea(r.Plateau.Area)
	case constants.DirectionSouth:
		r.Position.SubY()
		isInsideArea = r.Position.IsInsideArea(r.Plateau.Area)
	case constants.DirectionEast:
		r.Position.AddX()
		isInsideArea = r.Position.IsInsideArea(r.Plateau.Area)
	case constants.DirectionWest:
		r.Position.SubX()
		isInsideArea = r.Position.IsInsideArea(r.Plateau.Area)
	}

	if !isInsideArea {
		return exception.FellFromMars
	}

	return nil
}

func (r *Rover) turnRight() (err error) {
	switch r.Direction {
	case constants.DirectionNorth:
		r.Direction, err = valueObject.NewDirection(constants.DirectionEast)
		break
	case constants.DirectionEast:
		r.Direction, err = valueObject.NewDirection(constants.DirectionSouth)
		break
	case constants.DirectionSouth:
		r.Direction, err = valueObject.NewDirection(constants.DirectionWest)
		break
	case constants.DirectionWest:
		r.Direction, err = valueObject.NewDirection(constants.DirectionNorth)
		break
	}

	return err
}

func (r *Rover) turnLeft() (err error) {
	err = r.turnRight()
	if err != nil {
		return err
	}

	err = r.turnRight()
	if err != nil {
		return err
	}

	err = r.turnRight()
	if err != nil {
		return err
	}

	return nil
}

func (r *Rover) String() string {
	return fmt.Sprintf("%d %d %s", r.Position.GetX(), r.Position.GetY(), r.Direction)
}
