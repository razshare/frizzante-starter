package expired

import f "github.com/razshare/frizzante"

type Controller struct{}

func (_ Controller) Configure(meta func() f.Metadata) f.PageConfiguration {
	return f.PageConfiguration{
		Metadata: meta(),
	}
}

func (_ Controller) Base(req *f.Request, res *f.Response) {
	// Noop.
}

func (_ Controller) Action(req *f.Request, res *f.Response) {
	// Noop.
}
