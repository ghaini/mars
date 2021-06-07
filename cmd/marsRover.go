package cmd

import (
	"fmt"
	"log"
	"mars-rover/src/domain"
	"mars-rover/src/valueObject"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mars-rover",
	Short: "A client for move rovers on the mars",
	Run:   marsRoverHandler,
}

func marsRoverHandler(cmd *cobra.Command, args []string) {
	if len(args) < 6 || (len(args)-2)%4 != 0 {
		log.Fatalln("The number of submitted parameters is incorrect")
	}

	plateauX, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(err)
	}

	plateauY, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalln(err)
	}

	plateauArea := valueObject.NewPoint(plateauX, plateauY)
	plateau := domain.NewPlateau(plateauArea)

	for i := 2; i < len(args); i += 4 {
		positionX, err := strconv.Atoi(args[i])
		if err != nil {
			log.Fatalln(err)
		}

		positionY, err := strconv.Atoi(args[i+1])
		if err != nil {
			log.Fatalln(err)
		}

		direction, err := valueObject.NewDirection(args[i+2])
		if err != nil {
			log.Fatalln(err)
		}

		position := valueObject.NewPosition(positionX, positionY)
		rover, err := domain.NewRover(plateau, position, direction)
		if err != nil {
			log.Fatalln(err)
		}

		for j := 0; j < len(args[i+3]); j++ {
			move, err := valueObject.NewMove(string(args[i+3][j]))
			if err != nil {
				log.Fatalln(err)
			}

			err = rover.Move(move)
			if err != nil {
				log.Fatalln(err)
			}
		}

		err = plateau.AddRover(rover)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(rover)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
