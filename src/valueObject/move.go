package valueObject

import (
	"mars-rover/constants"
	"mars-rover/src/exception"
)

type Move string

func NewMove(move string) (Move, error) {
	switch move {
	case
		constants.MoveForward,
		constants.MoveRight,
		constants.MoveLeft:
		return Move(move), nil
	}
	return "", exception.InvalidMove
}
