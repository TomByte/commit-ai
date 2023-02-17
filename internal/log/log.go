package log

import (
	"github.com/fatih/color"
)

func Error(messages ...any) {
	c := color.New(color.FgRed)
	m := append([]any{"▶"}, messages...)
	c.Println(m...)
}

func Info(messages ...any) {
	c := color.New(color.FgCyan)
	m := append([]any{"▶"}, messages...)
	c.Println(m...)
}

func Success(messages ...any) {
	c := color.New(color.FgGreen)
	m := append([]any{"▶"}, messages...)
	c.Println(m...)
}
