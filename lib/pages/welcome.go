package pages

import f "github.com/razshare/frizzante"

func Welcome(
	withPath func(path string),
	withDocument func(doc *f.Document),
	_ func(baseHandler func(req *f.Request, res *f.Response, doc *f.Document)),
	_ func(actionFunction func(req *f.Request, res *f.Response, doc *f.Document)),
) {
	withPath("/")
	withDocument(f.DocumentCreate("welcome"))
}
