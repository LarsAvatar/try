package floatproblem

import (
	"fmt"
	"math"
	"strings"
)

type Cents int64

const maxIterationsFloat64 = 1_000_000_000 //This many iterations takes about 2 seconds on my laptop
var minIterationsFloat64 = maxIterationsFloat64

type Counter struct {
	Addend            Cents
	iterationsFloat32 int
	iterationsFloat64 int
}

func TryFloat64(pennies Cents) {
	amt := float64(pennies) / 100
	sumFloat := 0.00
	sumPennies := Cents(0)
	for x := 1; x < minIterationsFloat64; x++ {
		sumFloat += amt
		sumPennies += pennies
		sumFloatRounded := int(math.Round(sumFloat * 100))
		diff := sumFloatRounded - int(sumPennies)
		if diff != 0 {
			fmt.Printf(
				"Sum of Pennies = %s\nSum of float64 = %s\nAt iteration %d of adding %d pennies.\n\n",
				sumPennies.FmtComma(), FmtIntComma(sumFloatRounded), x, pennies)
			if x < minIterationsFloat64 {
				minIterationsFloat64 = x
			}
			break
		}
	}
}

func (amt Cents) Fmt() string {
	var out string
	if amt < 0 {
		out = fmt.Sprintf("%03d", amt)
	} else {
		out = fmt.Sprintf("%02d", amt)
	}
	sz := len(out)
	return out[:sz-2] + "." + out[sz-2:]
}

func (amt Cents) FmtComma() string {
	out := amt.Fmt()
	if (amt >= 0 && amt < 100000) || (amt <= 0 && amt > -100000) {
		return out
	}
	centsOnly := out[len(out)-3:]   //save the pennies
	dollarsOnly := out[:len(out)-3] //chop off the pennies
	negative := false
	if dollarsOnly[0] == '-' {
		negative = true
		dollarsOnly = dollarsOnly[1:]
	}
	var b strings.Builder
	if negative {
		b.WriteByte('-')
	}
	copyLen := len(dollarsOnly) % 3
	if copyLen > 0 {
		b.WriteString(dollarsOnly[:copyLen])
	}
	pos := copyLen
	for pos < len(dollarsOnly) {
		if copyLen > 0 {
			b.WriteByte(',')
		}
		copyLen = 3
		b.WriteString(dollarsOnly[pos : pos+copyLen])
		pos += copyLen
	}
	b.WriteString(centsOnly)
	return b.String()
}

func FmtFloat64(amt float64) string {
	return fmt.Sprintf("%.2f", amt)
}

func FmtFloat64Comma(amt float64) string {
	out := FmtFloat64(amt)
	if (amt >= 0 && amt < 100000) || (amt <= 0 && amt > -100000) {
		return out
	}
	centsOnly := out[len(out)-3:]   //save the pennies
	dollarsOnly := out[:len(out)-3] //chop off the pennies
	negative := false
	if dollarsOnly[0] == '-' {
		negative = true
		dollarsOnly = dollarsOnly[1:]
	}
	var b strings.Builder
	if negative {
		b.WriteByte('-')
	}
	copyLen := len(dollarsOnly) % 3
	if copyLen > 0 {
		b.WriteString(dollarsOnly[:copyLen])
	}
	pos := copyLen
	for pos < len(dollarsOnly) {
		if copyLen > 0 {
			b.WriteByte(',')
		}
		copyLen = 3
		b.WriteString(dollarsOnly[pos : pos+copyLen])
		pos += copyLen
	}
	b.WriteString(centsOnly)
	return b.String()
}

func FmtInt(amt int) string {
	var out string
	if amt < 0 {
		out = fmt.Sprintf("%03d", amt)
	} else {
		out = fmt.Sprintf("%02d", amt)
	}
	sz := len(out)
	return out[:sz-2] + "." + out[sz-2:]
}

// Format an integer amount with commas
func FmtIntComma(amt int) string {
	out := FmtInt(amt)
	if (amt >= 0 && amt < 100000) || (amt <= 0 && amt > -100000) {
		return out
	}
	centsOnly := out[len(out)-3:]   //save the pennies
	dollarsOnly := out[:len(out)-3] //chop off the pennies
	negative := false
	if dollarsOnly[0] == '-' {
		negative = true
		dollarsOnly = dollarsOnly[1:]
	}
	var b strings.Builder
	if negative {
		b.WriteByte('-')
	}
	copyLen := len(dollarsOnly) % 3
	if copyLen > 0 {
		b.WriteString(dollarsOnly[:copyLen])
	}
	pos := copyLen
	for pos < len(dollarsOnly) {
		if copyLen > 0 {
			b.WriteByte(',')
		}
		copyLen = 3
		b.WriteString(dollarsOnly[pos : pos+copyLen])
		pos += copyLen
	}
	b.WriteString(centsOnly)
	return b.String()
}
