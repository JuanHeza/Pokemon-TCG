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

var (
	win *pixelgl.Window
)

func main() {
	log.Println("Bienvenido al Pokemon TCG")
	//Battle("User", "Larry")
	Decode("NoCompile/Pokemon.json")
	pixelgl.Run(run)
}

func run() {
	var Icon []pixel.Picture
	I, err := loadPicture("Resource/Logo.png")
	Icon = append(Icon, I)
	if err != nil {
		panic(err)
	}
	cfg := pixelgl.WindowConfig{
		Title:     "Pokemon TCG",
		Bounds:    pixel.R(0, 0, 240, 160),
		VSync:     true,
		Resizable: true,
		Icon:      Icon,
	}
	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	BG, err := DrawBackground()
	if err != nil {
		panic(err)
	}
	log.Println(pixel.PictureDataFromImage(BG).Bounds(), "|", pixel.PictureDataFromImage(BG).Bounds().Center())
	Background := pixel.NewSprite(pixel.PictureDataFromImage(BG), pixel.R(0, 0, 240, 320))
	//root := pixel.V(120, 0)
	win.Clear(colornames.Skyblue)
	Background.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	BattleTest()
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

func loadPictureImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
