package logger

import (
	"fmt"
)

var (
	verbose bool
)

func Init(isVerbose bool) {
	verbose = isVerbose
}

func Debugln(format string, args ...any) {
	if verbose {
		fmt.Printf(format+"\n", args...)
	}
}

func Debug(format string, args ...any) {
	if verbose {
		fmt.Printf(format, args...)
	}
}
