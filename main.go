package main

import (
	"embed"
	. "github.com/razshare/frizzante"
)

//go:embed www/dist/*/**
var efs embed.FS

type Item struct {
	Checked     bool   `json:"checked"`
	Description string `json:"description"`
}

func main() {
	// Create.
	server := ServerCreate()
	var items = []Item{
		{Checked: false, Description: "Pet the cat."},
		{Checked: false, Description: "Do laundry"},
		{Checked: false, Description: "Pet the cat."},
		{Checked: false, Description: "Cook"},
		{Checked: false, Description: "Pet the cat."},
	}

	// Configure.
	ServerWithPort(server, 8080)
	ServerWithHostName(server, "127.0.0.1")
	ServerWithEmbeddedFileSystem(server, efs)
	var configure = func(_ *Request, _ *Response) *SveltePageConfiguration {
		return &SveltePageConfiguration{
			Render: ModeFull,
			Props: map[string]interface{}{
				"items": &items,
			},
		}
	}

	// Route.
	ServerWithSveltePage(server, "GET /", "welcome", configure)
	ServerWithSveltePage(server, "GET /todo", "todo", configure)
	ServerWithRequestHandler(server,
		"POST /check", func(server *Server, request *Request, response *Response) {
			if !VerifyContentType(request, "application/json") {
				SendStatus(response, 400)
				return
			}
			ReceiveJson(request, &items)
		})

	// Log.
	ServerOnError(server, func(err error) {
		ServerLogError(server, err)
	})
	ServerOnInformation(server, func(information string) {
		ServerLogInformation(server, information)
	})

	// Start.
	ServerStart(server)
}
