// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package display

func Example_structAsKey() {
	type Employee struct {
		Id   int
		Name string
	}

	salaries := map[Employee]int{
		{502131, "Yoshiki Shibata"}: 1701,
		{16690, "Yoshiki Shibata"}:  1701,
	}

	Display("salaries", salaries)
	// Unordered output:
	// Display salaries (map[display.Employee]int):
	// salaries[display.Employee{Id: 502131, Name: "Yoshiki Shibata"}] = 1701
	// salaries[display.Employee{Id: 16690, Name: "Yoshiki Shibata"}] = 1701
}

func Example_arrayAsKey() {
	name := [2]string{"Yoshiki", "Shibata"}

	salaries := map[[2]string]int{
		name: 1701,
	}

	Display("salaries", salaries)
	// Output:
	// Display salaries (map[[2]string]int):
	// salaries[[2]string{"Yoshiki", "Shibata"}] = 1701
}
