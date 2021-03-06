package main

import (
	"bingbong/item"
	"strings"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

func main() {
	vecty.RenderBody(&App{})
}

type App struct {
	vecty.Core
	state item.Item
}

func (c *App) body() *vecty.HTML {
	if c.state.Body == "" {
		c.state.Body = "body_base"
	}
	return img(c.state.Body)
}

func (c *App) eyes() *vecty.HTML {
	if c.state.Eyes == "" {
		c.state.Eyes = "eyes_base"
	}
	return img(c.state.Eyes)
}

func (c *App) mouth() *vecty.HTML {
	if c.state.Mouth == "" {
		c.state.Mouth = "mouth_base"
	}
	return img(c.state.Mouth)
}

func (c *App) glasses() *vecty.HTML {
	if c.state.Glasses == "" {
		c.state.Glasses = "empty"
	}
	return img(c.state.Glasses)
}

func (c *App) hat() *vecty.HTML {
	if c.state.Hat == "" {
		c.state.Hat = "empty"
	}
	return img(c.state.Hat)
}

func img(s string) *vecty.HTML {
	return elem.Image(
		vecty.Markup(
			prop.Src("/assets/"+s+".png"),
			vecty.Style("position", "absolute"),
			vecty.Style("top", "0px"),
			vecty.Style("width", "100%"),
		),
	)
}

func (c *App) Render() vecty.ComponentOrHTML {
	vecty.SetTitle("Générateur de BingBong")
	return elem.Body(
		vecty.Markup(vecty.Style("padding-top", "30px")),
		elem.Div(
			vecty.Markup(vecty.Class("container")),
			elem.Div(
				vecty.Markup(vecty.Class("row")),
				elem.Div(
					vecty.Markup(vecty.Class("col-8")),
					c.body(),
					c.eyes(),
					c.mouth(),
					c.glasses(),
					c.hat(),
				),
				elem.Div(
					vecty.Markup(vecty.Class("col-4")),
					elem.Button(
						vecty.Markup(
							vecty.Class("btn", "btn-secondary"),
							event.Click(c.onReset),
						),
						vecty.Text("Reset"),
					),
					vecty.Text(" "),
					elem.Button(
						vecty.Markup(
							vecty.Class("btn", "btn-secondary"),
							event.Click(c.onRandom),
						),
						vecty.Text("Aléatoire"),
					),
					elem.Break(),
					elem.Break(),
					c.renderTools(),
				),
			),
		),
	)
}

func (c *App) onReset(*vecty.Event) {
	c.state.Body = ""
	c.state.Eyes = ""
	c.state.Glasses = ""
	c.state.Hat = ""
	c.state.Mouth = ""
	vecty.Rerender(c)
}

func (c *App) onItemClick(item string) func(*vecty.Event) {
	return func(*vecty.Event) {
		print(item)
		if strings.HasPrefix(item, "body") {
			c.state.Body = item
		} else if strings.HasPrefix(item, "mouth") {
			c.state.Mouth = item
		} else if strings.HasPrefix(item, "eyes") {
			c.state.Eyes = item
		} else if strings.HasPrefix(item, "hat") {
			c.state.Hat = item
		} else {
			c.state.Glasses = item
		}
		vecty.Rerender(c)
	}
}

func (c *App) renderTools() *vecty.HTML {
	return elem.Div(
		c.renderTool("Corps", item.Bodies),
		elem.Break(),
		c.renderTool("Yeux", item.Eyes),
		elem.Break(),
		c.renderTool("Bouches", item.Mouthes),
		elem.Break(),
		c.renderTool("Lunettes", item.Glasses),
		elem.Break(),
		c.renderTool("Chapeaux", item.Hats),
	)
}

func (c *App) renderTool(title string, items []string) *vecty.HTML {
	return elem.Div(
		vecty.Markup(vecty.Class("card")),
		elem.Div(
			vecty.Markup(vecty.Class("card-body")),
			elem.Heading5(
				vecty.Markup(vecty.Class("card-title")),
				vecty.Text(title),
			),
			c.renderItems(items),
		),
	)
}

func (c *App) renderItems(items []string) *vecty.HTML {
	var list vecty.List
	for _, item := range items {
		list = append(list,
			elem.Span(
				elem.Image(
					vecty.Markup(
						prop.Src("/assets/"+item+".png"),
						vecty.Style("width", "20%"),
						event.Click(c.onItemClick(item)),
					),
				),
			))
	}
	return elem.Div(list)
}

func (c *App) onRandom(*vecty.Event) {
	c.state = item.Random()
	vecty.Rerender(c)
}
