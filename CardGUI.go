package main

import(
	"os"
	"errors"
	"image"
	"image/draw"
	
	"github.com/faiface/pixel"
)

var(
/*
size		  0,  0, 86,120
pic			  6, 13, 79, 61 	// 74*49
title		  4,  4, 82, 11 	// 79* 8
data		  4, 63, 88, 99 	// 79*37
weakness	  4,106, 82,115		// 79*10
line		  5,101, 70,104		// 66* 4
*/
	CardSize				= image.Rect(  0,  0, 71, 99)
	Bounds = map[string]image.Rectangle{
	"StageCard" 			: image.Rect(  5,  5, 76,104), 	//1,1
	"BasicCard" 			: image.Rect( 80,  5,151,104),	//1,2
	"PsychicCard"	 		: image.Rect(155,  5,225,104),	//1,3
	"WaterCard" 			: image.Rect(230,  5,300,104),	//1,4
	"ColorlessCard" 		: image.Rect(  5,108, 75,207),	//2,1
	"DarknessClasicCard" 	: image.Rect( 80,108,150,207),	//2,2
	"DarknessModernCard" 	: image.Rect(155,108,225,207),	//2,3
	"DragonClasicCard" 		: image.Rect(230,108,300,207),	//2,4
	"DragonModernCard" 		: image.Rect(  5,211, 75,310),	//3,1
	"FireClasicCard" 		: image.Rect( 80,211,150,310),	//3,2
	"FireModernCard" 		: image.Rect(155,211,225,310),	//3,3
	"FighingCard" 			: image.Rect(230,211,300,310),	//3,4
	"GrassCard" 			: image.Rect(  5,314, 75,413),	//4,1
	"LightingCard" 			: image.Rect( 80,314,150,413),	//4,2
	"MetalModernCard"		: image.Rect(155,314,225,413),	//4,3
	"MetalClasicCard" 		: image.Rect(230,314,300,413),	//4,4
	//"FairyCard" 			: image.Rect(  0,  0,  0,  0)	//
	}
)
//DrawCard Visor of cards
func DrawCard(Stage string, Type string)(pixel.Picture, error){
	file, err := os.Open("Resource/PokemonCard.png")
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


	Frame,ok := Bounds[Stage]
	if !ok {
		return nil, errors.New("Undefined Stage")
	}
	Background,ok = Bounds[Type]
	if !ok{
		return nil, errors.New("Undefined Type")
	}
	//draw.Draw(dst, r, src, sr.Min, draw.Src)
	draw.Draw(imageB, imageB.Bounds(), img, Background.Min, draw.Src)
    draw.Draw(imageB, imageB.Bounds(), img, Frame.Min, draw.Over)
	return pixel.PictureDataFromImage(imageB), nil
}
