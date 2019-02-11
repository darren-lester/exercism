package robot

import (
	"fmt"
)

// definitions used in step 1

var Step1Robot struct {
	X, Y int
	Dir
}

type Dir int

// var _ fmt.Stringer = Dir(1729)

func (d Dir) String() string {
	return string(d)
}

const (
	N = iota
	E
	S
	W
)

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

func Left() {
	Step1Robot.Dir--
	if Step1Robot.Dir < 0 {
		Step1Robot.Dir = W
	}
}

// additional definitions used in step 2

type Command byte // valid values are 'R', 'L', 'A'
type RU int
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type Step2Robot struct {
	Dir
	Pos
}

func (robot *Step2Robot) Advance() {
	switch robot.Dir {
	case N:
		robot.Pos.Northing++
	case E:
		robot.Pos.Easting++
	case S:
		robot.Pos.Northing--
	case W:
		robot.Pos.Easting--
	}
}

func (robot *Step2Robot) Turn(command Command) {
	switch command {
	case 'R':
		robot.Dir = (robot.Dir + 1) % 4
	case 'L':
		robot.Dir--
		if robot.Dir < 0 {
			robot.Dir = W
		}
	}
}

type Action Command

func StartRobot(commands chan Command, actions chan Action) {
	defer close(actions)
	for command := range commands {
		actions <- Action(command)
	}
}

func Room(extent Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {
	defer close(rep)
	for action := range act {
		switch action {
		case 'R', 'L':
			robot.Turn(Command(action))
		case 'A':
			if !robotWillCrash(extent, robot) {
				robot.Advance()
			}
		}
	}

	rep <- robot
}

func robotWillCrash(extent Rect, robot Step2Robot) bool {
	return (robot.Dir == N && robot.Pos.Northing == extent.Max.Northing) ||
		(robot.Dir == E && robot.Pos.Easting == extent.Max.Easting) ||
		(robot.Dir == S && robot.Pos.Northing == extent.Min.Northing) ||
		robot.Dir == W && robot.Pos.Easting == extent.Min.Easting
}

// additional definition used in step 3

type Step3Robot struct {
	Name string
	Step2Robot
}

type Action3 struct {
	Command rune
	Name    string
}

var validActions = map[rune]bool{
	'R': true, //turn right
	'L': true, // turn left
	'A': true, // advance
	'D': true, // done
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	for _, command := range script {
		action <- Action3{command, name}
	}
	action <- Action3{'D', name}
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {
	defer func() {
		report <- robots
	}()

	isValid, msg := isSimValid(extent, robots)
	if !isValid {
		log <- msg
		return
	}

	terminatedRobots := make(map[string]bool)

	for true {
		if len(terminatedRobots) == len(robots) {
			return
		}

		a := <-action

		if terminatedRobots[a.Name] {
			continue
		}

		if !validActions[a.Command] {
			log <- fmt.Sprintf("invalid action: %v", a)
			terminatedRobots[a.Name] = true
			continue
		}

		var robotToCommand *Step3Robot
		for i := 0; i < len(robots); i++ {
			robot := &robots[i]
			if robot.Name == a.Name {
				robotToCommand = robot
			}
		}

		if robotToCommand == nil {
			log <- fmt.Sprintf("unknown robot: %s", a.Name)
			terminatedRobots[a.Name] = true
			continue
		}

		switch a.Command {
		case 'L', 'R':
			robotToCommand.Turn(Command(a.Command))
		case 'A':
			if robotWillCrash(extent, robotToCommand.Step2Robot) {
				log <- fmt.Sprintf("robot %s about to crash", robotToCommand.Name)
				continue
			}

			robotPositions := getRobotPositions(robots)
			if robotsWillCollide(robotPositions, *robotToCommand) {
				log <- "robots attempting to advance into another"
				continue
			}

			robotToCommand.Advance()
		case 'D':
			terminatedRobots[robotToCommand.Name] = true
		}
	}
}

func isSimValid(extent Rect, robots []Step3Robot) (bool, string) {
	robotNames := make(map[string]bool)
	for _, r := range robots {
		if r.Name == "" {
			return false, "robot with no name"
		}

		if robotNames[r.Name] {
			return false, fmt.Sprintf("two robots with name: %s", r.Name)
		}
		robotNames[r.Name] = true
	}

	for _, r := range robots {
		if isOutOfBounds(extent, r.Pos) {
			return false, fmt.Sprintf("robot %s is out of bounds", r.Name)
		}
	}

	robotPositions := make(map[Pos]bool)
	for _, r := range robots {
		if robotPositions[r.Pos] {
			return false, fmt.Sprintf("two robots at position %v", r.Pos)
		}
		robotPositions[r.Pos] = true
	}

	return true, ""
}

func getRobotPositions(robots []Step3Robot) map[Pos]bool {
	robotPositions := make(map[Pos]bool)
	for _, r := range robots {
		robotPositions[r.Pos] = true
	}
	return robotPositions
}

func isOutOfBounds(extent Rect, position Pos) bool {
	return position.Northing < extent.Min.Northing ||
		position.Northing > extent.Max.Northing ||
		position.Easting < extent.Min.Easting ||
		position.Easting > extent.Max.Easting
}

func robotsWillCollide(robotPositions map[Pos]bool, robot Step3Robot) bool {
	robot.Advance()
	return robotPositions[robot.Pos]
}
