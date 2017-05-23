package core

import (
	"fmt"
	"os"
	"github.com/Member1221/plutobot-go/core/color"
)

// LogText contains the different types of text in a log.
type LogText struct {
	Text      string
	Verbosity string
}

// LogInfo logs information level Info.
func LogInfo(text string, owner string) LogText {
	fmt.Printf("[Info::%s] %s\n", owner, text)
	return LogText{text, "info"}
}

// LogInfoG logs information level Info, with green text c:.
func LogInfoG(text string, owner string) LogText {
	c := color.New(color.FgGreen)
	c.Printf("[Info::%s] ", owner)
	fmt.Printf(text + "\n")
	return LogText{text, "info"}
}

// LogWarning logs warning level Info.
func LogWarning(text string, owner string) LogText {
	c := color.New(color.FgYellow)
	c.Printf("[Warning::%s] ", owner)
	fmt.Printf(text + "\n")
	return LogText{text, "warn"}
}

// LogError logs error level Info.
func LogError(text string, owner string) LogText {
	c := color.New(color.FgRed)
	c.Printf("[Error::%s] ", owner)
	fmt.Printf(text + "\n")
	return LogText{text, "err"}
}

// LogFatal logs fatal error level Info and terminates the program.
func LogFatal(text string, owner string, errorcode int) {
	c := color.New(color.FgRed).Add(color.Bold)
	c.Printf("[FATAL ERROR::%s] ", owner)
	fmt.Printf(text + "\n")
	os.Exit(errorcode)
}
