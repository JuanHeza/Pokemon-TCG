package main

//go build && Pokemon_TCG.exe

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	log.Println("Bienvenido al Pokemon TCG")
	//pixelgl.Run(run)
	//Battle("User", "Larry")
	Decode("NoCompile/Pokemon.json")
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pokemon TCG",
		Bounds: pixel.R(0, 0, 240, 160),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	pic, err := DrawCard("BasicCard", "ColorlessCard")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pixel.R(0, 0, 80, 120))

	win.Clear(colornames.Skyblue)

	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
		win.Update()
	}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
