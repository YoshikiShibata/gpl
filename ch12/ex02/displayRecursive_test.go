// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package display

func Example_recursive() {
	type Cycle struct {
		Value int
		Tail  *Cycle
	}

	var c Cycle
	c = Cycle{42, &c}
	Display("c", c)
	// Output:
	// Display c (display.Cycle):
	// c.Value = 42
	// (*c.Tail).Value = 42
	// (*(*c.Tail).Tail).Value = 42
	// (*(*(*c.Tail).Tail).Tail).Value = 42
	// (*(*(*(*c.Tail).Tail).Tail).Tail).Value = 42
	// (*(*(*(*(*c.Tail).Tail).Tail).Tail).Tail).Value = 42
	// (*(*(*(*(*(*c.Tail).Tail).Tail).Tail).Tail).Tail).Value = 42
}
