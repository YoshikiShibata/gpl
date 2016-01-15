package tempflag

import (
	"ch07/ex06/tempconv"
	"flag"
	"fmt"
)

// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ tempconv.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature")
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of flag variable.
// The flag argument must have a quantity and a unit. e.g., "100C".
func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

// *fahrenheit satisfies the flag.Value interface
type fahrenheitFlag struct{ tempconv.Fahrenheit }

func (f *fahrenheitFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "F", "°F":
		f.Fahrenheit = tempconv.Fahrenheit(value)
		return nil
	case "C", "°C":
		f.Fahrenheit = tempconv.CToF(tempconv.Celsius(value))
		return nil
	}
	return fmt.Errorf("invalid temperature")
}

// FahrenheitFlag defines a Fahrenheitflag with the specified name,
// default value, and usage, and returns the address of flag variable.
// The flag argument must have a quantity and a unit. e.g., "100F".
func FahrenheitFlag(name string, value tempconv.Fahrenheit, usage string) *tempconv.Fahrenheit {
	f := fahrenheitFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Fahrenheit
}
