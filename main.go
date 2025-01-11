package main

import (
	"embed"
	frz "github.com/razshare/frizzante"
)

//go:embed www/dist/*
var embeddedFileSystem embed.FS

func main() {
	// Create.
	server := frz.ServerCreate()

	// Configure.
	frz.ServerWithPort(server, 8080)
	frz.ServerWithSecurePort(server, 8383)
	frz.ServerWithHostName(server, "127.0.0.1")
	frz.ServerWithEmbeddedFileSystem(server, embeddedFileSystem)
	frz.ServerWithTemporaryDirectory(server, ".temp")
	frz.ServerClearTemporaryDirectory(server)
	frz.ServerWithCertificateAndKey(server, "cert.pem", "key.pem")

	// Route.
	frz.ServerOnRequest(server, "GET /", func(Server *frz.Server, request *frz.Request, response *frz.Response) {
		if frz.RedirectToSecure(request, response) {
			return
		}

		frz.EmbeddedFileOrElse(request, response, func() {
			frz.FileOrElse(request, response, func() {
				frz.Svelte(response, frz.SvelteOptions{
					Ssr: true,
					Props: map[string]interface{}{
						"name": "world",
					},
				})
			})
		})
	})

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
