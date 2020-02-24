package main

import "github.com/veandco/go-sdl2/sdl"
import "github.com/veandco/go-sdl2/ttf"


type piece int

const ( //sets up goable varables
    blank piece = iota
    x
    o
    cat
)

type TicTacToeBoard struct { //sets up array type

	board []piece
}

func NewTicTacToeBoard() *TicTacToeBoard { //New game
    return &TicTacToeBoard{[]piece{blank,blank,blank,blank,blank,blank,blank,blank,blank}}
}

func (b *TicTacToeBoard) getWinner() int {//Chenks when win
	win := 0 //0 no one won //1 x won//2 o won//3 c won
	
		if(b.board[0] == x && b.board[1] == x && b.board[2] == x){ //X Wins
			win = 1
		}
		if(b.board[3] == x && b.board[4] == x && b.board[5] == x){ //X Wins
			win = 1
		}
		if(b.board[6] == x && b.board[7] == x && b.board[8] == x){ //X Wins
			win = 1
		}
		if(b.board[0] == x && b.board[3] == x && b.board[6] == x){ //X Wins
			win = 1
		}
		if(b.board[1] == x && b.board[4] == x && b.board[7] == x){ //X Wins
			win = 1
		}
		if(b.board[2] == x && b.board[5] == x && b.board[8] == x){ //X Wins
			win = 1
		}
		if(b.board[0] == x && b.board[4] == x && b.board[8] == x){ //X Wins
			win = 1
		}
		if(b.board[6] == x && b.board[4] == x && b.board[2] == x){ //X Wins
			win = 1
		}
		////////////////////////////////////////////////////////////////////
		
		if(b.board[0] == o && b.board[1] == o && b.board[2] == o){ //O Wins
			win = 2
		}
		if(b.board[3] == o && b.board[4] == o && b.board[5] == o){ //O Wins
			win = 2
		}
		if(b.board[6] == o && b.board[7] == o && b.board[8] == o){ //O Wins
			win = 2
		}
		if(b.board[0] == o && b.board[3] == o && b.board[6] == o){ //O Wins
			win = 2
		}
		if(b.board[1] == o && b.board[4] == o && b.board[7] == o){ //O Wins
			win = 2
		}
		if(b.board[2] == o && b.board[5] == o && b.board[8] == o){ //O Wins
			win = 2
		}
		if(b.board[0] == o && b.board[4] == o && b.board[8] == o){ //O Wins
			win = 2
		}
		if(b.board[6] == o && b.board[4] == o && b.board[2] == o){ //O Wins
			win = 2
		}
		
    return win
}

func (b *TicTacToeBoard) getPlayer() piece {//keeps track of player turn and when the board is filled out 
    c := 9 
    for _ , v :=  range b.board {
        if v == blank {
            c--
        }
    }
    if c%2 == 0 {
        return x
    }
    return o
}

var font *ttf.Font


func (b *TicTacToeBoard) printPieces(renderer *sdl.Renderer) {// Keeps track of placement 

	placementX := int32(80)
	placementY := int32(60)
	i:=0
	
	textSurf,_ := font.RenderUTF8Solid("X", sdl.Color{0,255,0,255})
	
	textTexture,_ := renderer.CreateTextureFromSurface(textSurf)
	for _,v :=  range b.board {
		
		if v == blank{
		}
		
        if v == x {
			textSurf,_ = font.RenderUTF8Solid("X", sdl.Color{0,255,0,255})
			textTexture,_ = renderer.CreateTextureFromSurface(textSurf)
			renderer.Copy(textTexture, nil, &sdl.Rect{placementX,placementY,textSurf.W/2,textSurf.H/2})//Shady x
            
        } else if v == o {
			textSurf,_ = font.RenderUTF8Solid("O", sdl.Color{0,255,0,255})
			textTexture,_ = renderer.CreateTextureFromSurface(textSurf)
			renderer.Copy(textTexture, nil, &sdl.Rect{placementX,placementY,textSurf.W/2,textSurf.H/2})//Place x
		}
		i++
		if i % 3 > 0{
			placementX = placementX + 200
			
        }else if i % 3 == 0{
			placementY = placementY + 200
			placementX = 70
         
        } 
	} 
}


func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	
	// Using the SDL_ttf library so need to initialize it before using it
	if err := ttf.Init(); err != nil {
		panic(err)
	}
	defer ttf.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		600, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	renderer,err := sdl.CreateSoftwareRenderer(surface)
	
	if err != nil {
		panic(err)
	}
	
	font,err=ttf.OpenFont("Ubuntu-B.ttf",100)
	if err != nil {
		panic(err)
	}

	running := true
	lastx,lasty:=int32(0),int32(0)
	mouseDown:=false
	
	rect := &sdl.Rect{0,0,600,600}//setting the board
	color := uint32(0x000000)
	
	textSurf,_ := font.RenderUTF8Solid("O", sdl.Color{0,150,0,255})
	textTexture,_ := renderer.CreateTextureFromSurface(textSurf)
	
	var win int = 0
	var victory int = 0
	
	for running {
		board:=NewTicTacToeBoard()
		victory = 0
		for i := 0; i < 10; {
			rect = &sdl.Rect{0,0,600,600}//setting the board
			color = uint32(0x000000)
			surface.FillRect(rect, color)
		
			rect = &sdl.Rect{190,0,20,600}
			color = uint32(0x509950)
			surface.FillRect(rect, color)
		
			rect = &sdl.Rect{390,0,20,600}
			surface.FillRect(rect, color)
		
			rect = &sdl.Rect{0, 190, 600, 20}
			surface.FillRect(rect, color)
				
			rect = &sdl.Rect{0,390, 600, 20}
			surface.FillRect(rect, color)
			position:=lastx/200+lasty/200*3
			
			if i % 2 == 1 {
				textSurf,_ = font.RenderUTF8Solid("O", sdl.Color{0,150,0,255})
				textTexture,_ = renderer.CreateTextureFromSurface(textSurf)
				if board.board[position]==blank {
					renderer.Copy(textTexture, nil, &sdl.Rect{(lastx/200*200)+80,(lasty/200*200)+60,textSurf.W/2,textSurf.H/2})//Shady x
				}
		
				if mouseDown {
					textSurf,_ = font.RenderUTF8Solid("O", sdl.Color{0,0,255,255})
					textTexture,_ = renderer.CreateTextureFromSurface(textSurf)
					if board.board[position] == blank {
						renderer.Copy(textTexture, nil, &sdl.Rect{(lastx/200*200)+80,(lasty/200*200)+60,textSurf.W/2,textSurf.H/2})//Place x
				}
			}
			}else{
				textSurf,_ = font.RenderUTF8Solid("X", sdl.Color{0,150,0,255})
				textTexture,_ = renderer.CreateTextureFromSurface(textSurf)
				if board.board[position] == blank {
					renderer.Copy(textTexture, nil, &sdl.Rect{(lastx/200*200)+80,(lasty/200*200)+60,textSurf.W/2,textSurf.H/2})//Shady x
				}
		
				if mouseDown {
					textSurf,_ = font.RenderUTF8Solid("X", sdl.Color{0,0,255,255})
					textTexture,_ = renderer.CreateTextureFromSurface(textSurf)
					if board.board[position] == blank {
					renderer.Copy(textTexture, nil, &sdl.Rect{(lastx/200*200)+80,(lasty/200*200)+60,textSurf.W/2,textSurf.H/2})//Place x
				}
				}
			}
			
			board.printPieces(renderer)
				
			if victory == 1{
				surface.FillRect(&sdl.Rect{130,180,330,240}, uint32(0x5099950))
				surface.FillRect(&sdl.Rect{140,190,310,220}, uint32(0x000000))
				
				textSurf,_ = font.RenderUTF8Solid("X Wins!", sdl.Color{0,250,0,255})
				textTexture,_ = renderer.CreateTextureFromSurface(textSurf)
				renderer.Copy(textTexture, nil, &sdl.Rect{200,250,textSurf.W/2,textSurf.H/2})//X WIIIINNNNNNNNNNNNNNNNNNNNNNNNNNNSSSSSSSSSSSSSSSSSSSSSAAAAAAAAA
			}
			if victory == 2{
				surface.FillRect(&sdl.Rect{130,180,330,240}, uint32(0x5099950))
				surface.FillRect(&sdl.Rect{140,190,310,220}, uint32(0x000000))
				
				textSurf,_ = font.RenderUTF8Solid("O Wins!", sdl.Color{0,250,0,255})
				textTexture,_ = renderer.CreateTextureFromSurface(textSurf)
				renderer.Copy(textTexture, nil, &sdl.Rect{200,250,textSurf.W/2,textSurf.H/2})//O WINS
			}
			if victory == 3{
				surface.FillRect(&sdl.Rect{130,180,330,240}, uint32(0x5099950))
				surface.FillRect(&sdl.Rect{140,190,310,220}, uint32(0x000000))
				
				textSurf,_ = font.RenderUTF8Solid("Cat Steals", sdl.Color{0,250,0,255})
				textTexture,_ = renderer.CreateTextureFromSurface(textSurf)
				renderer.Copy(textTexture, nil, &sdl.Rect{180,220,textSurf.W/2,textSurf.H/2})//At the end of the day, we are all losers
				
				textSurf,_ = font.RenderUTF8Solid("The ", sdl.Color{0,250,0,255})
				textTexture,_ = renderer.CreateTextureFromSurface(textSurf)
				renderer.Copy(textTexture, nil, &sdl.Rect{250,270,textSurf.W/2,textSurf.H/2})
				
				textSurf,_ = font.RenderUTF8Solid("Competition ", sdl.Color{0,250,0,255})
				textTexture,_ = renderer.CreateTextureFromSurface(textSurf)
				renderer.Copy(textTexture, nil, &sdl.Rect{150,310,textSurf.W/2,textSurf.H/2})
			}
			
				
			window.UpdateSurface()
	
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			window.UpdateSurface()
			switch e := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.MouseMotionEvent:
				lastx,lasty=e.X,e.Y
			case *sdl.MouseButtonEvent:
				if e.Button==sdl.BUTTON_LEFT && e.State==sdl.PRESSED {
					mouseDown=true
				}
				if e.Button==sdl.BUTTON_LEFT && e.State==sdl.RELEASED {
					mouseDown=false

						
					if (e.X < 200 && e.Y < 200 && board.board[0] == blank){
						board.board[0] = board.getPlayer()
						i++
					}
					if (e.X > 200 && e.X < 400 && e.Y < 200 && board.board[1] == blank){
						board.board[1] = board.getPlayer()
						i++
					}
					if (e.X > 400 && e.Y < 200 && board.board[2] == blank){
						board.board[2] = board.getPlayer()
						i++
					}
					if (e.X < 200 && e.Y > 200 && e.Y < 400 && board.board[3] == blank){
						board.board[3] = board.getPlayer()
						i++
					}
					if (e.X > 200 && e.X < 400 && e.Y > 200 && e.Y < 400 && board.board[4] == blank){
						board.board[4] = board.getPlayer()
						i++
					}
					if (e.X > 400 && e.Y > 200 && e.Y < 400 && board.board[5] == blank){
						board.board[5] = board.getPlayer()
						i++
					}
					if (e.X < 200 && e.Y > 400 && board.board[6] == blank){
						board.board[6] = board.getPlayer()
						i++
					}
					if (e.X > 200 && e.X < 400 && e.Y > 400 && board.board[7] == blank){
						board.board[7] = board.getPlayer()
						i++
					}
					if (e.X > 400 && e.Y > 400 && board.board[8] == blank){
						board.board[8] = board.getPlayer()
						i++
					}
					if(e.X > 0 && e.Y > 0 && victory == 3 || victory == 2 || victory == 1){
					i++
					}	
				}
			} 
		}
		win = board.getWinner()
			
				if win == 1 && victory == 0{
					i = 9					
					victory = 1
						
				
				}
				if win == 2 && victory == 0{
					i = 9
					victory = 2
				
				}
				if win == 0 && i == 9 {
					i = 9
					victory = 3
				
			}	
		}
	}
			
}		




