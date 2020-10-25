package main

import "github.com/LarsAvatarUMG/try/floatproblem"

func main() {
	for pennies := 21400; pennies < 21499; pennies++ {
		floatproblem.TryFloat64(floatproblem.Cents(pennies))
	}
}
