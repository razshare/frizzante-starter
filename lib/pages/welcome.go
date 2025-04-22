package pages

import f "github.com/razshare/frizzante"

func Welcome(
	withPath func(path string),
	withDocument func(document *f.Document),
	_ func(baseHandler func(request *f.Request, response *f.Response, document *f.Document)),
	_ func(actionFunction func(request *f.Request, response *f.Response, document *f.Document)),
) {
	withPath("/")
	withDocument(f.DocumentCreate("Welcome"))
}
