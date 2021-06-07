package exception

import "fmt"

var (
	AlreadyOccupied  = fmt.Errorf("This position has already been occupied by another rover")
	IsNotOnMars      = fmt.Errorf("The rover is not on Mars")
	FellFromMars     = fmt.Errorf("The rover fell from Mars")
	InvalidMove      = fmt.Errorf("Invalid Move")
	InvalidDirection = fmt.Errorf("Invalid Direction")
)
