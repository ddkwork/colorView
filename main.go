package main

import (
	"colorView"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("colorView", func(w *unison.Window) {
		w.Content().AddChild(colorView.New().Layout())
	})
}
