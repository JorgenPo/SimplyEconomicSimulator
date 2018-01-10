package main

import (
	"fmt"
	"time"
	"os"
	"github/jorgen/hello/simulator"
	"bufio"
	"github/jorgen/hello/util"
)

func welcomeScreen() {
	logo := [26]string{
		" _______ _________ _______  ________ _",
		"(  ____ \\\\__   __/(       )(  ____ )( \\   |\\     /|",
		"| (    \\/   ) (   | () () || (    )|| (   ( \\   / )",
		"| (_____    | |   | || || || (____)|| |    \\ (_) /",
		"(_____  )   | |   | |(_)| ||  _____)| |     \\   /",
		"      ) |   | |   | |   | || (      | |      ) (",
		"/\\____) |___) (___| )   ( || )      | (____/\\| |",
		"\\_______)\\_______/|/     \\||/       (_______/\\_/",
		" ",
		" _______  _______  _______  _        _______  _______ _________ _______",
		"(  ____ \\(  ____ \\(  ___  )( (    /|(  ___  )(       )\\__   __/(  ____ \\",
		"| (    \\/| (    \\/| (   ) ||  \\  ( || (   ) || () () |   ) (   | (    \\/",
		"| (__    | |      | |   | ||   \\ | || |   | || || || |   | |   | |",
		"|  __)   | |      | |   | || (\\ \\) || |   | || |(_)| |   | |   | |",
		"| (      | |      | |   | || | \\   || |   | || |   | |   | |   | |",
		"| (____/\\| (____/\\| (___) || )  \\  || (___) || )   ( |___) (___| (____/\\",
		"(_______/(_______/(_______)|/    )_)(_______)|/     \\|\\_______/(_______/",
		"",
		" _______           _______           _        _______ _________ _______  _______",
		"(  ____ \\|\\     /|(       )|\\     /|( \\      (  ___  )\\__   __/(  ___  )(  ____ )",
		"| (    \\/| )   ( || () () || )   ( || (      | (   ) |   ) (   | (   ) || (    )|",
		"| (_____ | |   | || || || || |   | || |      | (___) |   | |   | |   | || (____)|",
		"(_____  )| |   | || |(_)| || |   | || |      |  ___  |   | |   | |   | ||     __)",
		"      ) || |   | || |   | || |   | || |      | (   ) |   | |   | |   | || (\\ (  ",
		"/\\____) || (___) || )   ( || (___) || (____/\\| )   ( |   | |   | (___) || ) \\ \\__       v0.1a",
		"\\_______)(_______)|/     \\|(_______)(_______/|/     \\|   )_(   (_______)|/   \\__/       By JorgenPo!",
	}

	for i := 0; i < 26; i++ {
		util.ClearScreen()
		for j := 0; j <= i; j++ {
			fmt.Println(logo[j])
		}

		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}

func mainMenuScreen() (int, error) {
	currentChoise := 0
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("== Use first letters of menu entries to navigate. Hit enter to choose menu ==")

		if currentChoise == 0 {
			fmt.Printf("=>")
		}
		fmt.Printf("New game\n")

		if simulator.HasSavedGames() {
			if currentChoise == 1 {
				fmt.Printf("=>")
			}
			fmt.Printf("Load game\n")
		}

		if currentChoise == 2 {
			fmt.Printf("=>")
		}
		fmt.Printf("Settings\n")

		if currentChoise == 3 {
			fmt.Printf("=>")
		}
		fmt.Printf("Credits\n")

		if currentChoise == 4 {
			fmt.Printf("=>")
		}
		fmt.Printf("Exit\n")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("ERROR OCCURED!\n")
			return -1, err
		}

		switch input[0] {
		case 'N':
			currentChoise = 0
		case 'L':
			if simulator.HasSavedGames() {
				currentChoise = 1
			}
		case 'S':
			currentChoise = 2
		case 'C':
			currentChoise = 3
		case 'E':
			currentChoise = 4
		case '\n':
			util.ClearScreen()
			return currentChoise, nil
		}

		util.ClearScreen()
	}

	return currentChoise, nil
}

func main() {
	welcomeScreen()

	fmt.Println()
	fmt.Println("Press any key to continue...")
	fmt.Scanln()
	util.ClearScreen()

	for {
		menuIndex, err := mainMenuScreen()

		if err != nil {
			fmt.Printf("Internal error: %v\n", err)
			fmt.Println("Game will be exited! Sorry. We will handle this in future versions!")
			time.Sleep(time.Duration(5000))
			os.Exit(1)
		}

		simulator.RunAction(menuIndex)
	}
}
