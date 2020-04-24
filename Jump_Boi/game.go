package main

import ."./gameEngine"
//import "github.com/veandco/go-sdl2/sdl"
//import "github.com/veandco/go-sdl2/gfx"
//import "github.com/veandco/go-sdl2/ttf"

func main(){
	Setup();
	s:=NewCustomSprite(280,-400,30,50,0,0)
	AddSprite(s)
	Run()
}
