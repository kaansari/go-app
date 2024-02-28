package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type homePage struct {
	app.Compo
}

func newHomePage() *homePage {
	return &homePage{}
}

func (p *homePage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *homePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *homePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
	analytics.Page("home", nil)
}

func (p *homePage) Render() app.UI {
	return newPage().
		Title("NOON").
		Icon("/web/icon.png").
		Content(
			ui.Flow().
				StretchItems().
				Spacing(84).
				Content(
					newRemoteMarkdownDoc().
						Class("fill").
						Src("/web/documents/homepage.md"),
					newRemoteMarkdownDoc().
						Class("fill").
						Class("updates").
						Src("/web/documents/updates.md"),
				),
			app.Div().Class("separator"),

			newRemoteMarkdownDoc().Src("/web/documents/home.md"),

			app.Div().Class("separator"),

			newBuiltWithGoapp().ID("built-with-go-app"),

			app.Div().Class("separator"),

			newRemoteMarkdownDoc().Src("/web/documents/home-next.md"),
		)

}
