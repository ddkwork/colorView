package main

import (
	"colorView"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("colorView", func(w *unison.Window) {
		w.Content().AddChild(colorView.New().Layout())
	})
}
