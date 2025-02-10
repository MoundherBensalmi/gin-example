package log_colors

import (
	"log"
	"os"
)

var Reset = "\033[0m"
var Green = "\033[32m"
var GreenIcon = "✅"

var Yellow = "\033[33m"
var YellowIcon = "⚠️"

var Red = "\033[31m"
var RedIcon = "❌"

var Blue = "\033[34m"
var BlueIcon = "🔄"

func CLog(color rune, message ...any) {
	usedColor := Reset
	switch color {
	case 'g':
		usedColor = Green + GreenIcon
	case 'y':
		usedColor = Yellow + YellowIcon
	case 'r':
		usedColor = Red + RedIcon
	case 'b':
		usedColor = Blue + BlueIcon
	}

	logArgs := append([]any{usedColor}, message...)
	logArgs = append(logArgs, Reset)
	log.Println(logArgs...)
}

func CFLog(color rune, message ...any) {
	CLog(color, message...)
	os.Exit(1)
}
