package main

import . "github.com/razshare/frizzante"

func main() {
	// Setup.
	server := ServerCreate()
	ServerWithPort(server, 8080)

	// Logging.
	ServerOnResponseError(server, func(request *Request, response *Response, err error) {
		println("Response Error -", err.Error())
	})
	ServerOnError(server, func(err error) {
		println("Error -", err.Error())
	})
	ServerOnInformation(server, func(information string) {
		println("Information -", information)
	})

	// Workspace.
	workspace := WorkspaceCreate()
	WorkspaceWithTemporaryDirectory(workspace, "temp")
	WorkspaceWithNodeModulesDirectory(workspace, "node_modules")
	render, compileError := WorkspaceCompileSvelte(workspace, "index.svelte")
	if compileError != nil {
		println(compileError.Error())
		return
	}

	// Route.
	ServerOnRequest(server, "GET", "/", func(request *Request, response *Response) {
		html, renderError := render(map[string]any{})
		if renderError != nil {
			ServerNotifyResponseError(server, request, response, renderError)
			return
		}
		Echo(response, html)
	})

	// Start.
	startError := ServerStart(server)
	if startError != nil {
		println(startError.Error())
		return
	}
}
