package main

import (
	getir "github.com/alimgiray/getir/internal"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	g := new(getir.Getir)
	g.Initialize()
	g.Run()
}
