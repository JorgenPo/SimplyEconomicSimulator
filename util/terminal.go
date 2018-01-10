package util

import (
	"runtime"
	"os/exec"
	"os"
	"fmt"
	"time"
	"bufio"
)

func ClearScreen() {
	if runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		return
	}

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Wait(miliseconds int) {
	time.Sleep(time.Duration(miliseconds) * time.Millisecond)
}

func PrintWithDefaultDelay(text string) {
	PrintWithDelay(text, 600)
}

func PrintWithDelay(text string, delay int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println(text)
}

func PrintAndWaitInput(text string) {
	fmt.Println(text)
	WaitInput()
}

func WaitInput() {
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func TextSlideDown(lines []string, lineSpace int) {
	for step := 0; step < len(lines); step++ {
		ClearScreen()
		for i := 0; i <= step; i++ {
			fmt.Println(lines[step - i])
			for j := 0; j < lineSpace; j++ {
				fmt.Println()
			}
		}

		Wait(1200)
	}
}

