package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type architecturePage struct {
	app.Compo
}

func newArchitecturePage() *architecturePage {
	return &architecturePage{}
}

func (p *architecturePage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *architecturePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *architecturePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Understanding go-app Architecture")
	ctx.Page().SetDescription("Documentation about how go-app parts are working together to form a Progressive Web App (PWA).")
	analytics.Page("architecture", nil)
}

func (p *architecturePage) Render() app.UI {
	return app.Div().Body(

		newRemoteMarkdownDoc().Src("/web/documents/architecture.md"),
	)
}
