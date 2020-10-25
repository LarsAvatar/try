package main

import "github.com/LarsAvatarUMG/try/floatproblem"

// Find an amount that's off by a penny after adding it about a million times.
func main() {
	for pennies := 21400; pennies < 21499; pennies++ {
		floatproblem.TryFloat64(floatproblem.Cents(pennies))
	}
}
