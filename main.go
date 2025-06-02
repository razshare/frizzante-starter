package main

import (
	"embed"
	"github.com/razshare/frizzante"
	"log"
	"path/filepath"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	assets := frizzante.NewAssets()

	loadError := assets.LoadViews(filepath.Join("lib", "components", "views"))
	if loadError != nil {
		log.Fatal(loadError)
	}

	utilitiesError := assets.CreateUtilities(filepath.Join(".frz", "utilities"))
	if utilitiesError != nil {
		log.Fatal(utilitiesError)
	}

	routerError := assets.CreateRouter(filepath.Join(".frz", "router"))
	if routerError != nil {
		log.Fatal(routerError)
	}

	//frz.NewServer().
	//	WithDist(dist).
	//	AddRoute(frz.Route{Pattern: "GET /", Handler: handlers.GetDefault}).
	//	AddRoute(frz.Route{Pattern: "GET /welcome", Handler: handlers.GetWelcome}).
	//	AddRoute(frz.Route{Pattern: "GET /todos", Handler: handlers.GetTodos}).
	//	AddRoute(frz.Route{Pattern: "POST /todos", Handler: handlers.PostTodos}).
	//	Start()
}
