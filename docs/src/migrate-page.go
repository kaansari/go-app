//aboutus

package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type migratePage struct {
	app.Compo
}

func newMigratePage() *migratePage {
	return &migratePage{}
}

func (p *migratePage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *migratePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *migratePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("About Us")
	ctx.Page().SetDescription("Our Story")
	analytics.Page("aboutus", nil)
}

func (p *migratePage) Render() app.UI {
	return newPage().
		Title("About Us").
		//Icon(swapSVG).
		Content(
			newRemoteMarkdownDoc().Src("/web/documents/aboutus.md"),
		)
}



