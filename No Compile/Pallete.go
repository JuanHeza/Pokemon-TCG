package main
//23.10.07
//23.10.14
//23.
import (
    "fmt"
	"image/color"
	_"os"
	_"image/jpeg"
	"image"
	_"log"
)

var Paleta15Bits []color.Color
var Paleta color.Palette

func Pallete(){
	//a, b ,c := 31, 31<<5, 31<<10
	// fmt.Printf(" len16(%016b) - %d\n",a,a)
	// fmt.Printf(" len16(%016b) - %d\n",b,b)
	// fmt.Printf(" len16(%016b) - %d\n",c,c)
	//valor acumulado * 8
	for R:= 0; R <32; R++{
		for G:= 0; G <32; G++{
			for B:= 0; B <32; B++{
				//RGB := R<<10 + G <<5 + B
				Paleta15Bits = append(Paleta15Bits,color.RGBA{uint8(R*8),uint8(G*8),uint8(B*8),uint8(255)})
			}
		}
	}
	Paleta = Paleta15Bits
	fmt.Println("Paleta Iniciada")
	//Convert("001.jpg")
}

func Convert(input image.Image)(image.Image){//(input string){
	/*file, err := os.Open(input)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}*/
		img := input
		b := img.Bounds()
		imgSet := image.NewRGBA(b)
		for X := 0; X < b.Max.X; X++ {
			for Y := 0; Y < b.Max.Y; Y++ {
				oldPixel := img.At(X, Y)
				pixel := Paleta.Convert(oldPixel)
				imgSet.Set(X,Y,pixel)
			}
		}
		/*
		outputfile, err := os.Create("Out.jpg")
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)*/
		return imgSet
}