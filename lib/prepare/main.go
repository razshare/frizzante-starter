package main

import (
	frz "github.com/razshare/frizzante"
	"path"
)

func main() {
	frz.PrepareStart()
	frz.PrepareSveltePages(path.Join("lib", "pages"))
	frz.PrepareEnd()
}
