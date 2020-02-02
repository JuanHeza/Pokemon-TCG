package main 
//go run pallete.go utility.go
import (
	"log"
	"bytes"
	"golang.org/x/net/html"
	"net/http"
	"io/ioutil"
	"errors"
	"io"
	"strings"
	"image"
	"os"
	"image/jpeg"
	"github.com/nfnt/resize"
	"image/draw"
)

var(
	store []*html.Node
	Base = "https://www.serebii.net"
	Pic []image.Image
	Sets = map[string]string{
		"jungle/" : "https://www.serebii.net/card/jungle",
		"fossil/" : "https://www.serebii.net/card/fossil",
		"teamrocket/" : "https://www.serebii.net/card/teamrocket",
		"gymheroes/" : "https://www.serebii.net/card/gymheroes",
//		"gymchallenge/" : "https://www.serebii.net/card/gymchallenge",
	}
)

func main() {
	for CardSet,Page := range Sets{
		store = []*html.Node{}
		Pic = []image.Image{}
		rsp, err := http.Get(Page)
		if err != nil {
			log.Fatal(err) 
		}
		defer rsp.Body.Close()

		buf, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			log.Fatal(err) 
		}
		str1 := string(buf)

		doc, _ := html.Parse(strings.NewReader(str1))
		bn, err := GetTag(doc,"table",false)
		if err != nil {
			return
		}
		_,err = GetTag(bn,"tr",true)
		for i,b := range store{
			if i != 0{
				dat := downloadImage(GetImage(b,i,CardSet))
				if dat != nil{
					Pic = append(Pic,dat)
				}
			}
			log.Println("#5",len(Pic),i,"\n")
		}
		Pallete()
		CreateSheet(Pic,strings.Trim(CardSet,"/"))
		log.Println("done",CardSet)
	}
}

func GetTag(doc *html.Node, tag string, storing bool) (*html.Node, error) {
	var i = 0
	var body *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == tag {
			body = node
			if storing{
				store = append(store,body)
			}
			i++
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if body != nil {
		return body, nil
	}
	return nil, errors.New("Missing <body> in the node tree")
}
  
func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func GetImage(input *html.Node, Serial int, CardSet string) (string, int){
	//var aux = false
	var body *html.Node
	for child := input.FirstChild; child != nil; child = child.NextSibling{
		for kid := child.FirstChild; kid != nil; kid = kid.NextSibling{
			for baby := kid.FirstChild; baby != nil; baby = baby.NextSibling{
				if baby.Type == html.ElementNode && baby.Data == "img"{
					for _, n := range baby.Attr{
						if n.Key == "src" && strings.Contains(n.Val,CardSet){
							body = baby
							for _,dat := range body.Attr{
								if dat.Key == "src"{
									return dat.Val,1
								}
							}
						}
					}
				}
			}
		}
	}
	/*	var crawler func(*html.Node)
		crawler = func(node *html.Node){
			if node.Type == html.ElementNode && node.Data == "img" {
				if aux{
					body = node
					aux = false
					log.Println("#2.5", body.Attr, body.Attr[0])
					return 
				}else{
					aux = true
				}
			}
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				crawler(child)
			}
		}
		crawler(input)
		if body != nil {
			for _,dat := range body.Attr{
				log.Println("#2.7", dat.Key)
				if dat.Key == "src"{
					return dat.Val,1
				}
			}
		}
	*/	return "/error",0
}

func downloadImage(input string, val int)(image.Image){
	var img image.Image
	log.Println("#3",Base+input)
	if val != 0{
		client := http.Client{
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}
		resp, err := client.Get(Base+input)
		defer resp.Body.Close()
		img, err = jpeg.Decode(resp.Body)
		if err != nil {
			img, _ , err = image.Decode(resp.Body)
			if err != nil {
				return nil
			}
		}
		return img
	}
	return nil
}

func CreateSheet(Pics []image.Image, name string){
	var Pos = image.Rect(0,0,74,49)
	var X,Y int
	sheet := image.NewRGBA(image.Rect(0,0,(12*75),((len(Pics)/12)+1)*50))
	log.Println("Sheet",sheet.Bounds())
	for a, pic := range Pics{
		X = a / 12
		Y = a % 12
		m := resize.Resize(74,49, pic, resize.Lanczos3)
		n := Convert(m)
		draw.Draw(sheet, Pos, n, image.ZP, draw.Over)
		log.Print("#",a,"[",X,"] [",Y,"]",Pos)
		if Y == 11 {
			Pos = Pos.Add(image.Point{-825,50})
		}else{
			Pos = Pos.Add(image.Point{75,0})
		}
		if !Pos.In(sheet.Bounds()){
			log.Println("Impoible",Pos)
			os.Exit(0)
		}
		//Pos = image.Rect((X*74)+1,(Y*49)+1,((1+X)*74),((1+Y)*49))
		log.Println("done")
	}
	log.Println("Creating sheet")
	
	outputfile, err := os.Create(name+".jpg")
	if err != nil {
		log.Fatal(err)
	}
		defer outputfile.Close()
	jpeg.Encode(outputfile,sheet,nil)
}