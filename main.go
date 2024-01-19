package main

import (
	"fmt"
	//"log"
	//"combat.go"
	//"image/png"
	//"time"
	//"math"
	"github.com/hajimehoshi/ebiten/v2" //i wouldve tried working on my own engine but i didnt have the time
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//"github.com/hajimehoshi/ebiten/v2/vector"
	//"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	CLEAR = "\033[H\033[2J"
)

var (
	p = fmt.Print
	s = fmt.Scanln

	main_menu_check        bool
	intro_check            bool = false
	sidescrollingcheck     bool = true
	combatscript           bool
	capableofexiting       bool
	capableofsaving        bool
	startup                bool
	newsave                bool = true //hopefully i can store this on some online database to actually have game saves be functional
	ignoreintro_on_newgame bool

	characterchoice string

	choice1lvl int
	choice2lvl int
	choice3lvl int

	choice1hp int
	choice2hp int
	choice3hp int

	countdown  int
	countdown2 int

	item_id  int
	enemy_id int

	userinputfloat float64
	gamestatus     string
)

func main() {
	//	ebiten.SetWindowSize(screenX, screenY)
	ebiten.SetWindowTitle("KESSLER INDUCED OUROBOROS")
	//	initialselection()
	//	p(PM_N_PD_20SUNBEAM_44)
	//	intro()
	//thingamajiglol()
	//main1()

	//charactermove(&Game{})
	gamestatus = "a"
	decode()
}

/*
func main() {
	newsave=true
	s(newsave)
	switch newsave {
		case true:
		introcutscene()
		s(ignoreintro_on_newgame)
			switch ignoreintro_on_newgame {
				case true:
				intro_check=false
				case false:
				intro_check=true
			}
		case false:
		introcutscene()
		intro_check=false
	}
}*/
/*
func introcutscene () {
	//intro_check
	s(intro_check,sidescrollingcheck,combatscript,main_menu_check)
	switch (intro_check && sidescrollingcheck && combatscript && main_menu_check) {
		case intro_check==false:
		INTRO()
		case intro_check==true:
		MENU()
	}
//main menu script goes here, contemplating whether or not to have intro start before main menu or after player chooses to start a new game
/*	if main_menu_check==true {
		p("Bruh!")
	}*/
/*
}

func INTRO() {
	p("WHY DOESNT THIS RUN")
}

func MENU() {
	p("HUH")
	s(&userinputfloat)
	userinputfloat[]
	sidescrolling(userinputfloat[0],userinputfloat[1],userinputfloat[2])
}

strings.ToLower(userinput)

*/
///engine test, simple HELLO WORLD print
/*
type game struct{}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main1() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&game{}); err != nil {
		log.Fatal(err)
	}
}
*/
