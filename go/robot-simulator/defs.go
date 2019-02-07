package robot

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
