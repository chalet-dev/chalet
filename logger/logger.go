/*
 * Copyright © 2024 Ignacio Chalub <ignaciochalub@gmail.com> & Federico Pochat <federicopochat@gmail.com>
 */

package logger

import (
	"fmt"
	"github.com/fatih/color"
)

func Info(message string) {
	c := color.New(color.FgCyan)
	_, _ = c.Println(fmt.Sprintf("INFO: %s", message))
}

// Warn logs a warning message
func Warn(message string) {
	c := color.New(color.FgYellow)
	_, _ = c.Println(fmt.Sprintf("WARN: %s", message))
}

// Error logs an error message
func Error(message string) {
	c := color.New(color.FgRed)
	_, _ = c.Println(fmt.Sprintf("ERROR: %s", message))
}

// Print as white
func Print(message string) {
	fmt.Println(message)
}
