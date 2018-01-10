package simulator

import (
	"github/jorgen/hello/util"
	"os"
)

var actions = map[int]func() {
	0: newGame,
	1: loadGame,
	2: settings,
	3: credits,
	4: exit,
}

func RunAction(actionIndex int) {
	actions[actionIndex]()
	util.ClearScreen()
}

func newGame() {

}

func loadGame() {

}

func settings() {

}

func credits() {
	credits := []string {
		"Simply economic simulator game",
		"Lead Programmer: JorgenPO",
		"Sound Effects: JorgenPO",
		"Graphics: JorgenPO",
		"Idea: Revolusha",
		"Write your emails to: popoff96@live.com",
	}

	util.TextSlideDown(credits, 4)
	util.WaitInput()
}

func exit() {
	util.PrintAndWaitInput("Exiting Simply Economic Simulator! Goodbye!")
	os.Exit(0)
}