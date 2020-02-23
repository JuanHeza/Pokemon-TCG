package main

import (
	"image"
	"image/draw"
	"os"
)

var (
	//Background1 is the First layer, the background color in the battlefield the stadium
	Background1 image.Image
	//Background2 is the Second layer, the battlefield layer
	Background2 image.Image
	//Background3 is the Third layer, the Card position and energies atched
	Background3 image.Image
	//Background4 is the Fourth layer, Extra info layer, like names and so
	Background4 image.Image
	stadium     = "Resource/metal_classic.png"
	layer       = "Resource/mapBattleground.png"
	spaces      = map[string]image.Rectangle{
		//"default":	  image.Rect(0,0,29,39),
		"P1-Deck":    image.Rect(206, 252, 235, 291),
		"P1-Discard": image.Rect(206, 205, 235, 244),
		"P1-Trainer": image.Rect(178, 213, 207, 252),
		"P1-Main":    image.Rect(0, -41, 142, 252), //image.Rect(113, 213, 142, 252),
		"P1-MainB":   image.Rect(108, 218, 147, 247),
		"P1-Bench1":  image.Rect(180, 264, 209, 303),
		"P1-Bench2":  image.Rect(147, 264, 176, 303),
		"P1-Bench3":  image.Rect(114, 264, 143, 303),
		"P1-Bench4":  image.Rect(81, 264, 110, 303),
		"P1-Bench5":  image.Rect(48, 264, 77, 303),
		"P1-Price1":  image.Rect(3, 170, 32, 209),
		"P1-Price2":  image.Rect(8, 167, 37, 206),
		"P1-Price3":  image.Rect(3, 211, 32, 250),
		"P1-Price4":  image.Rect(8, 208, 37, 247),
		"P1-Price5":  image.Rect(3, 252, 32, 291),
		"P1-Price6":  image.Rect(8, 249, 37, 288),

		"P2-Deck":    image.Rect(4, 69, 33, 108),
		"P2-Discard": image.Rect(4, 116, 33, 155),
		"P2-Trainer": image.Rect(48, 132, 77, 171),
		"P2-Main":    image.Rect(0, 40, 0,0), //image.Rect(113, 132, 142, 171),
		"P2-MainB":   image.Rect(108, 137, 147, 166),
		"P2-Bench1":  image.Rect(46, 81, 75, 120),
		"P2-Bench2":  image.Rect(79, 81, 108, 120),
		"P2-Bench3":  image.Rect(112, 81, 141, 120),
		"P2-Bench4":  image.Rect(145, 81, 174, 120),
		"P2-Bench5":  image.Rect(178, 81, 207, 120),
		"P2-Price1":  image.Rect(207, 151, 236, 190),
		"P2-Price2":  image.Rect(203, 154, 231, 193),
		"P2-Price3":  image.Rect(207, 110, 236, 149),
		"P2-Price4":  image.Rect(203, 113, 231, 152),
		"P2-Price5":  image.Rect(207, 69, 236, 108),
		"P2-Price6":  image.Rect(203, 72, 231, 111),
	}
)

//Position is a struct to save the data in each of the slots in the battlefield/hand
type Position struct {
	ID    string
	Space image.Rectangle
	Img   image.Image
}

//DrawBackground Merge the four background layers in one image
func DrawBackground() (image.Image, error) {
	Background := image.NewRGBA(image.Rect(0, 0, 240, 320))
	layer1, err := os.Open(stadium)
	if err != nil {
		return nil, err
	}
	defer layer1.Close()
	img1, _, err := image.Decode(layer1)
	if err != nil {
		return nil, err
	}
	layer2, err := os.Open(layer)
	if err != nil {
		return nil, err
	}
	defer layer2.Close()
	img2, _, err := image.Decode(layer2)
	if err != nil {
		return nil, err
	}

	//draw.Draw(dst, r, src, sr.Min, draw.Src)
	draw.Draw(Background, Background.Bounds(), img1, image.Point{}, draw.Src)
	draw.Draw(Background, Background.Bounds(), img2, image.Point{}, draw.Over)
	return Background, nil
}
