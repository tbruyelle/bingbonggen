package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

func main() {
	vecty.RenderBody(&App{})
}

type App struct {
	vecty.Core
}

func staticPath(s string) string {
	return "file:///home/thomas/src/bingbong/" + s
}

func img(s string) *vecty.HTML {
	return elem.Image(
		vecty.Markup(
			prop.Src(staticPath("assets/"+s+".png")),
			vecty.Style("position", "absolute"),
			vecty.Style("top", "0px"),
			vecty.Style("width", "100%"),
		),
	)
}

func (c *App) Render() vecty.ComponentOrHTML {
	vecty.SetTitle("Générateur de BingBong")
	return elem.Body(
		elem.Div(
			vecty.Markup(vecty.Class("container")),
			elem.Div(
				vecty.Markup(vecty.Class("row")),
				elem.Div(
					vecty.Markup(vecty.Class("col-md-8")),
					img("body_base"),
					img("mouth_base"),
					img("eyes_base"),
				),
				elem.Div(
					vecty.Markup(vecty.Class("col-md-4")),
					vecty.Text("Outils"),
				),
			),
		),
	)
}
