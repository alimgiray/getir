package getir

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Getir struct {
	port   string
	server http.Server
}

func (g *Getir) Initialize() {
	g.port = os.Getenv("PORT")
	if g.port == "" {
		g.port = "3000"
	}
}

func (g *Getir) Run() {
	g.server = http.Server{
		Addr:    fmt.Sprintf(":%s", g.port),
		Handler: NewRouter(),
	}
	log.Fatal(g.server.ListenAndServe())
}
