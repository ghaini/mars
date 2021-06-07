package valueObject

import (
	"mars-rover/constants"
	"mars-rover/src/exception"
)

type Direction string

func NewDirection(direction string) (Direction, error) {
	switch direction {
	case
		constants.DirectionEast,
		constants.DirectionNorth,
		constants.DirectionWest,
		constants.DirectionSouth:
		return Direction(direction), nil
	}
	return "", exception.InvalidDirection
}
