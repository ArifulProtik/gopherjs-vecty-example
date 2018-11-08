package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

func main() {
	vecty.SetTitle("Simple Login Page With Vecty")
	vecty.AddStylesheet("test/style.css")
	vecty.RenderBody(&mybody{})
}

type mybody struct {
	vecty.Core
}

func (*mybody) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("mform"),
			),
			elem.Heading1(
				vecty.Text("Login"),
			),
			elem.Input(
				vecty.Markup(
					prop.Type(prop.TypeText),
					vecty.Class("inpt"),
					prop.Placeholder("Username"),
				),
			),
			elem.Input(
				vecty.Markup(
					prop.Type(prop.TypePassword),
					vecty.Class("inpt"),
					prop.Placeholder("password"),
				),
			),
			elem.Input(
				vecty.Markup(
					prop.Type(prop.TypeSubmit),
					prop.Value("Login"),
					vecty.Class("sbmt"),
				),
			),
			elem.Break(),
			elem.Span(
				vecty.Markup(
					vecty.Class("s1"),
				),
				vecty.Text("Not a member."),
				elem.Anchor(
					vecty.Markup(
						prop.Href("#"),
					),
					vecty.Text("Sign Up Here"),
				),
			),
		),
	)
}
