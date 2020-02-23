package main

import (
	"errors"
	"image"
	"image/draw"
	"log"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	/*
	   size		  0,  0, 86,120
	   pic			  6, 13, 79, 61 	// 74*49
	   title		  4,  4, 82, 11 	// 79* 8
	   data		  4, 63, 88, 99 	// 79*37
	   weakness	  4,106, 82,115		// 79*10
	   line		  5,101, 70,104		// 66* 4
	*/

	//CardSize is the frame of the card
	CardSize = image.Rect(0, 0, 71, 99)
	//Bounds is the map of the card resource sheet
	Bounds = map[string]image.Rectangle{
		"StageCard":          image.Rect(5, 5, 76, 104),      //1,1
		"BasicCard":          image.Rect(80, 5, 151, 104),    //1,2
		"PsychicCard":        image.Rect(155, 5, 225, 104),   //1,3
		"WaterCard":          image.Rect(230, 5, 300, 104),   //1,4
		"ColorlessCard":      image.Rect(5, 108, 75, 207),    //2,1
		"DarknessClasicCard": image.Rect(80, 108, 150, 207),  //2,2
		"DarknessModernCard": image.Rect(155, 108, 225, 207), //2,3
		"DragonClasicCard":   image.Rect(230, 108, 300, 207), //2,4
		"DragonModernCard":   image.Rect(5, 211, 75, 310),    //3,1
		"FireClasicCard":     image.Rect(80, 211, 150, 310),  //3,2
		"FireModernCard":     image.Rect(155, 211, 225, 310), //3,3
		"FighingCard":        image.Rect(230, 211, 300, 310), //3,4
		"GrassCard":          image.Rect(5, 314, 75, 413),    //4,1
		"LightingCard":       image.Rect(80, 314, 150, 413),  //4,2
		"MetalModernCard":    image.Rect(155, 314, 225, 413), //4,3
		"MetalClasicCard":    image.Rect(230, 314, 300, 413), //4,4
		//"FairyCard" 			: image.Rect(  0,  0,  0,  0)	//
	}

	miniBounds = map[Element]int{
		typeDarkness:  2,
		typeColorless: 3,
		typeFire:      4,
		typeGrass:     5,
		typeFighting:  6,
		typeDragon:    7,
		typeLightning: 8,
		typeMetal:     9,
		typePsychic:   10,
		typeWater:     11,
		//typeFairy:          image.Rect(5, 314, 75, 413),    //4,1
	}
)

//DrawCard Visor of cards
func DrawCard(Stage string, Type string) (pixel.Picture, error) {
	file, err := os.Open("Resource/MiniTexture.png")
	var Frame, Background image.Rectangle
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	imageB := image.NewRGBA(CardSize)

	Frame, ok := Bounds[Stage]
	if !ok {
		return nil, errors.New("Undefined Stage")
	}
	Background, ok = Bounds[Type]
	if !ok {
		return nil, errors.New("Undefined Type")
	}
	//draw.Draw(dst, r, src, sr.Min, draw.Src)
	draw.Draw(imageB, imageB.Bounds(), img, Background.Min, draw.Src)
	draw.Draw(imageB, imageB.Bounds(), img, Frame.Min, draw.Over)
	return pixel.PictureDataFromImage(imageB), nil
}

func (P CardPokemon) drawPKMiniCard() *image.RGBA {
	var Card = image.NewRGBA(image.Rect(0, 0, 30, 40))
	var maskPic = image.NewRGBA(image.Rect(0, 0, 30, 40))
	var maskType = image.NewRGBA(image.Rect(0, 0, 30, 40))
	var Px = image.Point{X: 30, Y: 0}
	var Py = image.Point{Y: 40, X: 0}
	var Size = image.Rect(0, 0, 30, 40)
	//	var pic = image.Rect(1, 7, 28, 29)
	//	var title = image.Rect(1, 1, 28, 5)
	//	var sub = image.Rect(1, 31, 28, 38)
	//	var weak = image.Rect(17, 33, 28, 38)
	//	var resi = image.Rect(3, 33, 6, 36)
	//drawing mask's
	border, err := loadPictureImage("Resource/Border layer_minimodel_1.png")
	checkError(err)
	draw.Draw(maskPic, Card.Bounds(), border, Px, draw.Src)
	draw.Draw(maskType, Card.Bounds(), border, image.ZP, draw.Src)
	for i := 0; i < miniBounds[P.Type]%4; i++ {
		Size = Size.Add(Px)
	}
	for j := 0; j < miniBounds[P.Type]/4; j++ {
		Size = Size.Add(Py)
	}
	element, err := loadPictureImage("Resource/minitexture.png")
	draw.Draw(Card, Card.Bounds(), element, Size.Min, draw.Src)
	//SetBorder
	draw.Draw(Card, Card.Bounds(), border, border.Bounds().Min, draw.Over)

	if len(P.Weaknesses.Cost) != 0 {
		waek, err := loadPictureImage("Resource/W layer_minimodel_1.png")
		checkError(err)
		draw.Draw(Card, Card.Bounds(), waek, image.ZP, draw.Over)
	}
	if len(P.Resistence.Cost) != 0 {
		resi, err := loadPictureImage("Resource/R layer_minimodel_1.png")
		checkError(err)
		draw.Draw(Card, Card.Bounds(), resi, image.ZP, draw.Over)
	}
	//2 mascaras, 11 tipos, 2 plantillas
	return Card
}

func drawSetPokemon(W *pixelgl.Window, P interface{}, pos string) {
	PImage := P.(CardPokemon).drawPKMiniCard()
	Ppixel := pixel.PictureDataFromImage(PImage)
	log.Println(Ppixel.Bounds(), "|", Ppixel.Bounds().Center())
	Pk1 := pixel.NewSprite(Ppixel, pixel.R(0, 0, 30, 40))
	X, Y := float64(spaces[pos].Min.X), float64(spaces[pos].Min.Y)
	Pk1.Draw(W, pixel.IM.Moved(W.Bounds().Center().Add(pixel.V(X, Y))))
}
