package main

import (
	"embed"
	frz "github.com/razshare/frizzante"
)

type Item struct {
	Checked     bool   `json:"checked"`
	Description string `json:"description"`
}

//go:embed www/dist/*/**
var efs embed.FS

var items = []Item{
	{Checked: false, Description: "Pet the cat."},
	{Checked: false, Description: "Do laundry"},
	{Checked: false, Description: "Pet the cat."},
	{Checked: false, Description: "Cook"},
	{Checked: false, Description: "Pet the cat."},
}

func configure(_ *frz.Request) *frz.SveltePageConfiguration {
	return &frz.SveltePageConfiguration{
		Render: frz.ModeFull,
		Props: map[string]interface{}{
			"items": &items,
		},
	}
}

func main() {
	// Create.
	server := frz.ServerCreate()

	// Setup.
	frz.ServerWithPort(server, 8080)
	frz.ServerWithHostName(server, "127.0.0.1")
	frz.ServerWithEmbeddedFileSystem(server, efs)

	// Route.
	frz.ServerWithSveltePage(server, "GET /", "welcome", configure)
	frz.ServerWithSveltePage(server, "GET /todo", "todo", configure)
	frz.ServerWithRequestHandler(server, "POST /check",
		func(server *frz.Server, request *frz.Request, response *frz.Response) {
			if !frz.VerifyContentType(request, "application/json") {
				frz.SendStatus(response, 400)
				return
			}
			frz.ReceiveJson(request, &items)
		},
	)

	// Log.
	frz.ServerOnError(server, func(err error) {
		frz.ServerLogError(server, err)
	})
	frz.ServerOnInformation(server, func(information string) {
		frz.ServerLogInformation(server, information)
	})

	// Start.
	frz.ServerStart(server)
}
