package indexes

import f "github.com/razshare/frizzante"

func Welcome(
	route func(path string, page string),
	_ func(showFunction func(req *f.Request, res *f.Response, p *f.Page)),
	_ func(actionFunction func(req *f.Request, res *f.Response, p *f.Page)),
) {
	route("/", "welcome")
}
