package main

type char struct {
	entity_id
	placeholdergui string
	name
	perfectability int
	reaction float64
	walkspeed float64
	leap float64
	backleap float64
	side2side float64
	recoilmanagement int
	ballistic float64
	autolead float64
	autotargetcone float64
	progress float64
	singleplayer bool
	multiplayer bool
	carrycapacity // in kilograms
}

var (

	//charactersatatime[] char
	DuskAmongstShatteredHorizons=char{1,"a","Shattered Horizons", 0, 37, 20, 4,9,15,1,1,.1,30,0,false,false,500}//bruh i didnt even intend for their acronym to be DASH
	DEBUGNOIR=char{0,"a","NOIR - DEBUG CHAR.",100,360,76.2,155,155,127,100,500,1,240,0,false, false,250000} //yea they arent supposed to be in this setting - too op bruh
)/*
func initialselection() {
	p("test")
}
*/