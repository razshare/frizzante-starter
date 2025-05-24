package expired

import f "github.com/razshare/frizzante"

type Controller struct{}

func (_ Controller) Configure(id f.Identifier) f.PageConfiguration {
	return f.PageConfiguration{
		Id: id(),
	}
}

func (_ Controller) Base(req *f.Request, res *f.Response) {
	// Noop.
}

func (_ Controller) Action(req *f.Request, res *f.Response) {
	// Noop.
}
