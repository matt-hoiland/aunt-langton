package ant

import "testing"

func TestString(t *testing.T) {
	tests := []struct {
		data     Direction
		expected string
	}{
		{Up, "Up"},
		{Down, "Down"},
		{Left, "Left"},
		{Right, "Right"},
		{-1, "Unknown"},
		{4, "Unknown"},
		{87, "Unknown"},
	}

	for _, test := range tests {
		if test.data.String() != test.expected {
			t.Errorf("Direction %d = %s, want %s", test.data, test.data, test.expected)
		}
	}
}

func TestTurnLeft(t *testing.T) {
	ant := Ant{i: 0, j: 0, direction: Up}

	ant.TurnLeft()
	if ant.direction != Left {
		t.Errorf("ant.direction = %s, want %s\n", ant.direction, Left)
	}
	ant.TurnLeft()
	if ant.direction != Down {
		t.Errorf("ant.direction = %s, want %s\n", ant.direction, Down)
	}
	ant.TurnLeft()
	if ant.direction != Right {
		t.Errorf("ant.direction = %s, want %s\n", ant.direction, Right)
	}
	ant.TurnLeft()
	if ant.direction != Up {
		t.Errorf("ant.direction = %s, want %s\n", ant.direction, Up)
	}
}

func TestTurnRight(t *testing.T) {
	ant := Ant{i: 0, j: 0, direction: Up}

	ant.TurnRight()
	if ant.direction != Right {
		t.Errorf("ant.direction = %s, want %s\n", ant.direction, Right)
	}
	ant.TurnRight()
	if ant.direction != Down {
		t.Errorf("ant.direction = %s, want %s\n", ant.direction, Down)
	}
	ant.TurnRight()
	if ant.direction != Left {
		t.Errorf("ant.direction = %s, want %s\n", ant.direction, Left)
	}
	ant.TurnRight()
	if ant.direction != Up {
		t.Errorf("ant.direction = %s, want %s\n", ant.direction, Up)
	}
}

func TestAdvance(t *testing.T) {
	tests := []struct {
		ant   Ant
		i, j  int
		steps int
	}{
		{Ant{100, 100, Up, 0}, 99, 100, 1},
		{Ant{100, 100, Right, 1}, 100, 101, 2},
		{Ant{100, 100, Down, 2}, 101, 100, 3},
		{Ant{100, 100, Left, 3}, 100, 99, 4},
	}

	for _, test := range tests {
		test.ant.Advance()
		if test.ant.i != test.i || test.ant.j != test.j {
			t.Errorf("Direction = %s, ant is at (%d, %d), should be at (%d, %d)\n", test.ant.direction, test.ant.i, test.ant.j, test.i, test.j)
		}
		if test.ant.steps != test.steps {
			t.Errorf("Ant took %d steps, should have taken %d steps\n", test.ant.steps, test.steps)
		}
	}
}
