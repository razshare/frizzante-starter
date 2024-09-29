package main

import . "github.com/razshare/frizzante"
import "path/filepath"

var temp, tempError = filepath.Abs(".temp")
var www, wwwError = filepath.Abs("www")
var nodeModules, nodeModulesError = filepath.Abs("node_modules")

func main() {
	// Check for errors.
	if nil != tempError {
		println(tempError.Error())
		return
	}
	if nil != wwwError {
		println(wwwError.Error())
		return
	}
	if nil != nodeModulesError {
		println(nodeModulesError.Error())
		return
	}

	// Setup.
	server := ServerCreate()
	ServerWithHostname(server, "127.0.0.1")
	ServerWithPort(server, 8080)
	ServerWithTemporaryDirectory(server, temp)
	ServerWithNodeModulesDirectory(server, nodeModules)
	ServerWithSvelteDirectory(server, "GET /", www)

	// Logging.
	ServerOnError(server, func(err error) {
		ServerLogError(server, err)
	})
	ServerOnInformation(server, func(information string) {
		ServerLogInformation(server, information)
	})

	// Start.
	startError := ServerStart(server)
	if startError != nil {
		println(startError.Error())
		return
	}
}
