package main

import (
	"math/rand"
	"strings"
	"time"

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

	state struct {
		body  string
		eyes  string
		mouth string
		bonus string
	}
}

func (c *App) body() *vecty.HTML {
	if c.state.body == "" {
		c.state.body = "body_base"
	}
	return img(c.state.body)
}

func (c *App) eyes() *vecty.HTML {
	if c.state.eyes == "" {
		c.state.eyes = "eyes_base"
	}
	return img(c.state.eyes)
}

func (c *App) mouth() *vecty.HTML {
	if c.state.mouth == "" {
		c.state.mouth = "mouth_base"
	}
	return img(c.state.mouth)
}

func (c *App) bonus() *vecty.HTML {
	if c.state.bonus == "" {
		c.state.bonus = "empty"
	}
	return img(c.state.bonus)
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
					vecty.Markup(vecty.Class("col-md-8")),
					c.body(),
					c.eyes(),
					c.mouth(),
					c.bonus(),
				),
				elem.Div(
					vecty.Markup(vecty.Class("col-md-4")),
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
	c.state.body = ""
	c.state.eyes = ""
	c.state.bonus = ""
	c.state.mouth = ""
	vecty.Rerender(c)
}

func (c *App) onItemClick(item string) func(*vecty.Event) {
	return func(*vecty.Event) {
		print(item)
		if strings.HasPrefix(item, "body") {
			c.state.body = item
		} else if strings.HasPrefix(item, "mouth") {
			c.state.mouth = item
		} else if strings.HasPrefix(item, "eyes") {
			c.state.eyes = item
		} else {
			c.state.bonus = item
		}
		vecty.Rerender(c)
	}
}

func (c *App) renderTools() *vecty.HTML {
	return elem.Div(
		c.renderTool("Corps", bodies),
		elem.Break(),
		c.renderTool("Yeux", eyes),
		elem.Break(),
		c.renderTool("Bouches", mouthes),
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
	rand.Seed(time.Now().UnixNano())
	c.state.body = bodies[rand.Intn(len(bodies))]
	c.state.mouth = mouthes[rand.Intn(len(mouthes))]
	c.state.eyes = eyes[rand.Intn(len(eyes))]
	vecty.Rerender(c)
}
