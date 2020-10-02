package ant

// Ant represents the ant which moves across the board
type Ant struct {
	i, j      int
	direction Direction
	steps     int
}

// NewAnt builds a new ant facing Up
func NewAnt(i, j int) Ant {
	return Ant{i, j, Up, 0}
}

// I gets the ant's i position
func (a Ant) I() int {
	return a.i
}

// J gets the ant's j position
func (a Ant) J() int {
	return a.j
}

// Direction gets the ant's current direction
func (a Ant) Direction() Direction {
	return a.direction
}

// Steps gets the ant's number of steps taken
func (a Ant) Steps() int {
	return a.steps
}

// Direction determines way the ant moves
type Direction int

// The set of 4 Directions
const (
	Right Direction = iota
	Up
	Left
	Down
)

func (d Direction) String() string {
	switch d {
	case Right:
		return "Right"
	case Up:
		return "Up"
	case Left:
		return "Left"
	case Down:
		return "Down"
	default:
		return "Unknown"
	}
}

// TurnLeft rotates direction of the Ant counterclockwise
func (a *Ant) TurnLeft() {
	a.direction++
	if a.direction > Down {
		a.direction = Right
	}
}

// TurnRight rotates the direction of the Ant clockwise
func (a *Ant) TurnRight() {
	a.direction--
	if a.direction < Right {
		a.direction = Down
	}
}

// Advance moves the ant 1 cell in the direction its facing
func (a *Ant) Advance() {
	switch a.direction {
	case Right:
		a.j++
	case Up:
		a.i--
	case Left:
		a.j--
	case Down:
		a.i++
	}
	a.steps++
}
