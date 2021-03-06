package domain

import (
	"log"
)

type Robot struct {
	*Postion
	Direction string
	Plane     *Postion
}

var direction = []string{"W", "N", "E", "S"}
var directionIndices = map[string]int{
	"W": 0,
	"N": 1,
	"E": 2,
	"S": 3,
}

func NewRobot(startingPostion, plane *Postion, direction string) *Robot {
	return &Robot{
		Postion:   startingPostion,
		Direction: direction,
		Plane:     plane,
	}
}

func (r *Robot) ProcessCommand(cmd string) {

	getDirection := func(dir string, turnClockWise bool) string {
		dirIndex := directionIndices[dir]
		if turnClockWise {
			if dirIndex == 3 {
				dirIndex = 0
			} else {
				dirIndex++
			}
		} else {
			if dirIndex == 0 {
				dirIndex = 3
			} else {
				dirIndex--
			}
		}
		return direction[dirIndex]
	}

	for _, c := range cmd {
		x, y := r.X, r.Y
		switch string(c) {
		case "L":
			r.Direction = getDirection(r.Direction, false)
		case "R":
			r.Direction = getDirection(r.Direction, true)
		case "M":
			switch r.Direction {
			case "W":
				x -= 1
			case "N":
				y += 1
			case "E":
				x += 1
			case "S":
				y -= 1
			}
		}

		if x > r.Plane.X || y > r.Plane.Y {
			log.Println("Command is not completely executed as the Robot goes out of plane.")
			break
		}
		r.X, r.Y = x, y
	}
}
