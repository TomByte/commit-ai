package log

import (
	"github.com/fatih/color"
)

func Error(messages ...any) {
	c := color.New(color.FgRed)
	c.Println(messages...)
}

func Info(messages ...any) {
	c := color.New(color.FgCyan)
	c.Println(messages...)
}

func Success(messages ...any) {
	c := color.New(color.FgGreen)
	c.Println(messages...)
}
