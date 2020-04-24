package gameEngine

import "github.com/veandco/go-sdl2/sdl"
import "github.com/veandco/go-sdl2/gfx"
import "github.com/veandco/go-sdl2/ttf"
import "math"
//import "fmt"
import "time"

type CustomSprite struct{
	 *Sprite
	 vx float64
	 vy float64
}

var Window *sdl.Window
var surface *sdl.Surface
var renderer *sdl.Renderer
var sprites []GameElement
var BGColor uint32 = 0xffffff
var keys []bool




type GameElement interface {
    GetX() float64
    GetY() float64
    GetW() int
    GetH() int
    KeyDown(int)
    KeyUp(int)
    Click(int,int)
    Act()
    Draw(*sdl.Renderer)
}
	
type Sprite struct {
    X float64
    Y float64
    W int
    H int
}

const (
	KeyLeft int = 80
	KeyRight int = 79
	KeyUp int = 82
	KeyDown int = 81
)

func (s *Sprite) GetX() float64 {return s.X}
func (s *Sprite) GetY() float64 {return s.Y}
func (s *Sprite) GetW() int {return s.W}
func (s *Sprite) GetH() int {return s.H}
func (s *Sprite) Click(x,y int) {}

func (s *CustomSprite) Act(){
	//s.X-=5
	w,_:=Window.GetSize()
	
	s.vx=0
	if IsKeyDown(KeyLeft) {
		s.vx-=5
	}
	if IsKeyDown(KeyRight) {
		s.vx+=5
	}
	s.X+=s.vx
	
	s.X+=float64(w)
	s.X=math.Mod(s.X,float64(w))
	
	s.Y+=s.vy
	if s.Y>500 {
		
		if(IsKeyDown(KeyDown)){
			s.Y=500
			s.vy=0
		}else{
			s.Y-=s.vy
			s.vy=-.9*s.vy
		}
		
	}
	s.vy+=.8 //GRAVITY
	
}

func (s *CustomSprite) KeyDown(x int) {
	if x==KeyUp && s.Y+1>500{
		s.vy=-25
	}
}

func (s *Sprite) KeyUp(k int) {}


func NewSprite(x,y float64,w,h int) *Sprite {
	return &Sprite{X:x,Y:y,W:w,H:h}
}

func (s *CustomSprite) Draw(r *sdl.Renderer){
	w,_:=Window.GetSize()
	r.SetDrawColor(0,0,255,255)
	if s.X>float64(int(w)-s.W) {
		r.FillRect(&sdl.Rect{int32(s.X-float64(w)), int32(s.Y), int32(s.W), int32(s.H)})
	}
	r.FillRect(&sdl.Rect{int32(s.X), int32(s.Y), int32(s.W), int32(s.H)})
}

func (s *Sprite) Draw(r *sdl.Renderer){
	gfx.FilledEllipseRGBA(r, int32(s.X+float64(s.W)/2), int32(s.Y+float64(s.H)/2),int32(s.W/2), int32(s.H/2), 0,0,0,255)
}

func NewCustomSprite(x,y float64,w,h int,vx,vy float64) *CustomSprite {
	return &CustomSprite{NewSprite(x,y,w,h),vx,vy}
}

func Close() {
	Window.Destroy()
	ttf.Quit()
	sdl.Quit()
}


func AddSprite(s GameElement) {
	sprites=append(sprites,s)
}

func Setup() {
	keys = make([]bool, 256)
	
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	if err := ttf.Init(); err != nil {
		panic(err)
	}
	
	var err error
	Window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		600, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	surface, err = Window.GetSurface()
	renderer,err = sdl.CreateSoftwareRenderer(surface)
	if err != nil {
		panic(err)
	}
}

func IsKeyDown(k int)bool{
	if(keys[k] == true){
		return true
	}else{
		return false
		}
	
}

func Run() {
	drawTick := time.Tick(16 * time.Millisecond)
	actTick := time.Tick(16 * time.Millisecond)
	for{
		select {
			case <-drawTick:
				DrawAll()
			case <-actTick:
				ActAll();
			//case <-exit:
			//	break
		}
		
	}
	Close()
	
}

func DrawAll(){
	surface.FillRect(nil, BGColor)
	for _,s:= range sprites {
		s.Draw(renderer)
	}
}

func ActAll() {
	
	copyOfSprites:=sprites
	for _,s:= range copyOfSprites { //clone sprites incase sprites kill each other
		s.Act()
	}
	
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				Close()
				return
			case *sdl.KeyboardEvent:
				if e.Type==sdl.KEYDOWN {
					if(!keys[int(e.Keysym.Scancode)]) {
						for _,s:= range sprites {
							s.KeyDown(int(e.Keysym.Scancode))
						}
					}
					keys[int(e.Keysym.Scancode)] = true
				}
				if e.Type==sdl.KEYUP {
					keys[int(e.Keysym.Scancode)] = false
					for _,s:= range sprites {
						s.KeyUp(int(e.Keysym.Scancode))
					}
				}
			case *sdl.MouseMotionEvent:
				
			case *sdl.MouseButtonEvent:
				
		}
	}
	Window.UpdateSurface()
}
