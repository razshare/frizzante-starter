package main

import . "github.com/razshare/frizzante"

func main() {
	// Setup.
	server := ServerCreate()
	ServerWithInterface(server, "127.0.0.1")
	ServerWithPort(server, 8080)

	// Logging.
	ServerOnError(server, func(err error) {
		println("Error -", err.Error())
	})
	ServerOnInformation(server, func(information string) {
		println("Information -", information)
	})

	// Workspace.
	workspace := WorkspaceCreate()
	WorkspaceWithTemporaryDirectory(workspace, ".temp")
	WorkspaceWithNodeModulesDirectory(workspace, "node_modules")
	render, compileError := WorkspaceCompileSvelte(workspace, "index.svelte")
	if compileError != nil {
		println(compileError.Error())
		return
	}

	// Route.
	ServerOnRequest(
		server, "GET", "/", func(server *Server, request *Request, response *Response) {
			html, renderError := render(map[string]any{
				"name": "world",
			})
			if nil != renderError {
				ServerNotifyError(server, renderError)
				return
			}
			Echo(response, html)
		},
	)

	// Start.
	startError := ServerStart(server)
	if startError != nil {
		println(startError.Error())
		return
	}
}
