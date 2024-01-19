package main

import (
	"log"
	//"bytes"
	"image"
	//"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	//"sprite/core/core_0007.png"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//"github.com/hajimehoshi/ebiten/v2/vector"
	// "github.com/hajimehoshi/ebiten/v2/inpututil"
	//_ "sprite/core"
)

/*
	type entity struct{
		entity_id
		FrontalRHAe[]int
		RearRHAe[]int
		PortRHAe[]int
		StarRHAe[]int
		velocity float64
		acceleration float64
		drag float64

//	pixel_loc
}
*/
type sprites struct {
	North, East, South, West *ebiten.Image
}

/*
	type sprite struct{
		spriteID int
		spritePrint string
		//pixel_loc
	}
*/

type tile struct {
	sprites []*ebiten.Image
}

type Game struct {
	currentLevel int ///*Level
	width        int
	height       int
	camX         x
	camY         y
	camScale     float64
	camScaleTo   float64
	mousePanX    x
	mousePanY    x
	px, py       int
	tiles        [][]*tile
	//offscreen    *ebiten.Image
}

/*ype input struct{
	non_numerical string
	numerical int
}*/
//store 3d array, space key for increase in elevation, backspace for decrease, a/d covers z axis, w/s covers horizontal movement (y and x axis hotkeys interchangeable depending on perspective, if i can implement that), along with coordinates, status, properties, interaction, and switching from movement to combat movement when necessary
//also have a map menu, quest/todolist, also on topic of quests, should we have a linear storyline or just have it be determined by user choice?idfkwhyitalktomyselflikethis
//idk how i would get diagonal movement to function
//nor how to properly animate the sprites as idk how to import and utilize my 3d assets into replit or even get them to function with golang

// /isometric bs idk why im overcomplicating shit
// type pixel_loc[][][]float64

type x float64
type y float64
type z float64

// (perspectivedistance/radius,d,d)d=dimension, x,y, alternatively z=1/d d=?
const (
	tilesize     int = 16 //
	screenwidth      = 1920
	screenheight     = 1080
)

/*
	type line struct {
		X1, Y1, X2, Y2 float64
	}
*/
var (
	screenX       = screenwidth
	screenY       = screenheight //assigned to a constant right now, if i get around to it i may be able to have the user configure their screen size
	tileY     int = 2 * tilesize //1ft/2px
	tileX     int = tilesize
	tileZ     int //just 16 per Z tile?
	tileYhalf = tileY / 2
	tileXhalf = tileX / 2
	core_0007 *ebiten.Image

//	keyboard_input=input{"W","A","S","D"," ","backspaceplaceholder","E","shiftplaceholder",1,2,3,4,5,6,7,8,9,0,}//input array to register space/backspace, w/s, a/d, interaction hotkeys, combat hotkeys, misc.
//
// pixel_loc[][][]float64
)

/*


x'=(x-z)/sqrt(2)
y'=(x+2y+z)/sqrt(6)



func thingamajiglol(pixel_loc[x][y][z]){
	s(&characterchoice)
	switch characterchoice {
		case "w":
		p("lol")
	}
}*/
/*
func sidescrolling(x float64, y float64, z float64) string{
	//placeholder
	x=1
	y=2
	z=3
	return x,y,z, //placeholder
}*/
/*
func movement(x float64, y float64, z float64){
		switch keyboard_inputs[3]{
			case "W": x+=x
			case "A": y+=y
			case "S": x-=x
			case "D": y-=y
		}
}*/

/*func spriteadd(North, East, South, West *ebiten.Image) *Sprite{
	return &Sprite{
		North, East, South, West,
	}
}*/

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//vector.DrawFilledRect(screen, float32(g.px)-2, float32(g.py)-2, 4, 4, color.Black, true)
	//vector.DrawFilledRect(screen, float32(g.px)-1, float32(g.py)-1, 2, 2, color.RGBA{255, 100, 100, 255}, true)

	// Split the image into horizontal lines and draw them with different scales.

	op := &ebiten.DrawImageOptions{}
	w, h := core_0007.Bounds().Dx(), core_0007.Bounds().Dy()
	for i := 0; i < h; i++ {
		op.GeoM.Reset()

		// Move the image's center to the upper-left corner.
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

		// Scale each lines and adjust the position.
		lineW := w + i*3/4
		x := -float64(lineW) / float64(w) / 2
		op.GeoM.Scale(float64(lineW)/float64(w), 1)
		op.GeoM.Translate(x, float64(i))

		// Move the image's center to the screen's center.
		op.GeoM.Translate(float64(screenX)/2, float64(screenY)/2)

		screen.DrawImage(core_0007.SubImage(image.Rect(0, i, w, i+1)).(*ebiten.Image), op)
	}
	//switch {

	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		//characterloc+=value
		//g.px -= 4
		p("debug a")
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		//characterloc+=value
		//g.px += 4
		p("debug d")
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		//characterloc+=value
		//g.py -= 4
		p("debug s")
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		//characterloc+=value
		//g.py += 4
		p("debug w")
	}

	//}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenX, screenY
}

func FogOfWar(FWx, FWy float64) {

}

/*
// input for character,,,,,,,,,,,,,,
func charactermove(g *Game) {
	/*
	   switch {
	   case ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA):

	   	//characterloc+=value
	   	//g.px -= 4
	   	p("debug a")

	   case ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD):

	   	//characterloc+=value
	   	//g.px += 4
	   	p("debug d")

	   case ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS):

	   	//characterloc+=value
	   	//g.py -= 4
	   	p("debug s")

	   case ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW):

	   		//characterloc+=value
	   		//g.py += 4
	   		p("debug w")
	   	}
*/ /*
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		//characterloc+=value
		//g.px -= 4
		p("debug a")
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		//characterloc+=value
		//g.px += 4
		p("debug d")
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		//characterloc+=value
		//g.py -= 4
		p("debug s")
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		//characterloc+=value
		//g.py += 4
		p("debug w")
	}
*/
//}

// finally
func decode() {
	img, _, err := ebitenutil.NewImageFromFile("sprite/core/core_0007.png")
	if err != nil {
		log.Fatal(err)
	}
	core_0007 = ebiten.NewImageFromImage(img)
	//p(core_0007)
	ebiten.SetWindowSize(screenX, screenY)
	//ebiten.SetWindowTitle("Perspective (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}

/*
	func rect(x, y, w, h float64) []line {
		return []line{
			{x, y, x, y + h},
			{x, y + h, x + w, y + h},
			{x + w, y + h, x + w, y},
			{x + w, y, x, y},
		}
	}
*/

func mainmenu() {
	img, _, err := ebitenutil.NewImageFromFile("sprite/core/core_0007.png")
	if err != nil {
		log.Fatal(err)
	}
	core_0007 = ebiten.NewImageFromImage(img)
	//p(core_0007)
	ebiten.SetWindowSize(screenX, screenY)
	//ebiten.SetWindowTitle("Perspective (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		p("DOWN")
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		p("UP")

	}
}
