package main

import(
    "log"
)

type File struct{

}


func (F *File)Decode(){
    log.Println("Decode")
}

func (F *File)Encode(){
    log.Println("Encode")
}