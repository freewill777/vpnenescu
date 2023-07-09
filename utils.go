package utils

import (
	"os"
	"os/exec"
)

func clear_console() {
	windows := exec.Command("cmd", "/c", "cls")
	windows.Stdout = os.Stdout
	windows.Run()

	linux := exec.Command("clear")
	linux.Stdout = os.Stdout
	linux.Run()
}
