package welcome

import f "github.com/razshare/frizzante"

func Index(
	withPage func(page string),
	withPath func(path string),
	_ func(baseHandler func(req *f.Request, res *f.Response, page *f.Page)),
	_ func(actionFunction func(req *f.Request, res *f.Response, page *f.Page)),
) {
	withPage("welcome")
	withPath("/")
}
