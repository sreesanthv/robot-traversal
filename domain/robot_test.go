package domain

import "testing"

func TestProcessCommand(t *testing.T) {
	robot := NewRobot(NewPostion(0, 0), NewPostion(4, 4), "N")
	robot.ProcessCommand("MMMRMMLM")

	if robot.Postion.X != 2 {
		t.Errorf("Invalid X position. Expected %d but got %d", 2, robot.Postion.X)
	}
	if robot.Postion.Y != 4 {
		t.Errorf("Invalid Y position. Expected %d but got %d", 4, robot.Postion.Y)
	}
	if robot.Direction != "N" {
		t.Errorf("Invalid direction. Expected %s but got %s", "N", robot.Direction)
	}
}
