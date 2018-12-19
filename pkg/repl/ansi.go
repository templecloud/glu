package repl

// Reference: http://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html 
// 

import (
	"fmt"
	"strings"
)

// ANSI escape codes.
const (
	// ANSI 8 Colour
	black   = "\u001b[30m"
	red     = "\u001b[31m"
	green   = "\u001b[32m"
	yellow  = "\u001b[33m"
	blue    = "\u001b[34m"
	magenta = "\u001b[35m"
	cyan    = "\u001b[36m"
	white   = "\u001b[37m"
	// ANSI Modifiers
	bright = ";1m"
	// Reset
	reset = "\u001b[0m"
)

// ANSI is a entity capable of converting the target into a string and
// applying the specified ANSI codes to the result.
type ANSI struct {
	Enabled bool
}

// NewANSI creates an ANSI entity.
func NewANSI(enabled bool) ANSI {
	return ANSI{Enabled: enabled}
}

func (a *ANSI) apply(ansi string, target interface{}) string {
	if a.Enabled {
		return fmt.Sprintf("%s%+v%s", ansi, target, reset)
	}
	return fmt.Sprintf("%+v", target)
}

// ANSI 8 Colour ==============================================================
//

func (a *ANSI) black(target interface{}) string {
	return a.apply(black, target)
}

func (a *ANSI) red(target interface{}) string {
	return a.apply(red, target)
}

func (a *ANSI) green(target interface{}) string {
	return a.apply(green, target)
}

func (a *ANSI) yellow(target interface{}) string {
	return a.apply(yellow, target)
}

func (a *ANSI) blue(target interface{}) string {
	return a.apply(blue, target)
}

func (a *ANSI) magenta(target interface{}) string {
	return a.apply(magenta, target)
}

func (a *ANSI) cyan(target interface{}) string {
	return a.apply(cyan, target)
}

func (a *ANSI) white(target interface{}) string {
	return a.apply(white, target)
}

// ANSI 16 Colour - Bright ====================================================
//

func (a *ANSI) brightBlack(target interface{}) string {
	return a.apply(strings.Replace(black, "m", bright, 1), target)
}

func (a *ANSI) brightRed(target interface{}) string {
	return a.apply(strings.Replace(red, "m", bright, 1), target)
}

func (a *ANSI) brightGreen(target interface{}) string {
	return a.apply(strings.Replace(green, "m", bright, 1), target)
}

func (a *ANSI) brightYellow(target interface{}) string {
	return a.apply(strings.Replace(yellow, "m", bright, 1), target)
}

func (a *ANSI) brightBlue(target interface{}) string {
	return a.apply(strings.Replace(blue, "m", bright, 1), target)
}

func (a *ANSI) brightMagenta(target interface{}) string {
	return a.apply(strings.Replace(magenta, "m", bright, 1), target)
}

func (a *ANSI) brightCyan(target interface{}) string {
	return a.apply(strings.Replace(cyan, "m", bright, 1), target)
}

func (a *ANSI) brightWhite(target interface{}) string {
	return a.apply(strings.Replace(white, "m", bright, 1), target)
}
