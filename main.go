package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/sreesanthv/robot-traversal/domain"
)

func main() {
	robot, command, err := readInputs()
	if err != nil {
		log.Fatal("Error in reading input:", err)
	}

	robot.ProcessCommand(command)
	fmt.Println("Robot position:", robot.Postion.X, robot.Postion.Y, robot.Direction)

}

func readInputs() (*domain.Robot, string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	readInput := func(label string) string {
		fmt.Print("\n", label)
		scanner.Scan()
		return strings.TrimSpace(scanner.Text())
	}

	isValidInput := func(check, pattern string) bool {
		valid, _ := regexp.Match(pattern, []byte(check))
		return valid
	}

	plane := readInput("Please enter top-right coordinates of the rectangular plan (X Y):")
	if !isValidInput(plane, `^\d+\s{1}\d+$`) {
		return nil, "", errors.New("Invalid top-right coordinates of the rectangular plan.")
	}

	roboStart := readInput("Please enter robot's starting position (X Y N):")
	if !isValidInput(roboStart, `^\d+\s{1}\d+\s{1}[NSEW]$`) {
		return nil, "", errors.New("Inalid robot starting position.")
	}

	command := readInput("Please enter command for robot:")
	if !isValidInput(command, `^[LRM]+$`) {
		return nil, "", errors.New("Inalid robot starting position.")
	}

	spt := strings.Split(plane, " ")
	x, _ := strconv.Atoi(spt[0])
	y, _ := strconv.Atoi(spt[1])
	planePos := domain.NewPostion(x, y)

	spt = strings.Split(roboStart, " ")
	x, _ = strconv.Atoi(spt[0])
	y, _ = strconv.Atoi(spt[1])
	d := spt[2]
	robot := domain.NewRobot(domain.NewPostion(x, y), planePos, d)

	return robot, command, nil
}
