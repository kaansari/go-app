package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type servicesPage struct {
	app.Compo
}

func newServicesPage() *servicesPage {
	return &servicesPage{}
}

func (p *servicesPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *servicesPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *servicesPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Services")
	ctx.Page().SetDescription("In the heart of Downtown Chicago with expertese in Botox, Dysport, PRP, Sculptra, facials, Cannula, Microneedling, IV Therapy at 231 S State St. Personalized, results-driven beauty")
	analytics.Page("prp", nil)
}

func (p *servicesPage) Render() app.UI {
	return newPage().
		Title("Services").
		Icon(rocketSVG).
		Index(
			newIndexLink().Title("Botox"),
			newIndexLink().Title("Fillers"),
			newIndexLink().Title("Breast PRP"),
			newIndexLink().Title("Facial PRP"),
			newIndexLink().Title("Microneedling"),
			app.Div().Class("separator"),

			newIndexLink().Title("Next"),
		).
		Content(
			newBuiltWithGoapp().ID("services"),
		)
}
