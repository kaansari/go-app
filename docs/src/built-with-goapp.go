package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type builtWithGoapp struct {
	app.Compo

	Iid    string
	Iclass string
}

func newBuiltWithGoapp() *builtWithGoapp {
	return &builtWithGoapp{}
}

func (b *builtWithGoapp) ID(v string) *builtWithGoapp {
	b.Iid = v
	return b
}

func (b *builtWithGoapp) Class(v string) *builtWithGoapp {
	b.Iclass = app.AppendClass(b.Iclass, v)
	return b
}

func (b *builtWithGoapp) Render() app.UI {
	return app.Div().
		Class(b.Iclass).
		Body(
			app.H2().
				ID(b.Iid).
				Text("Build With go-app"),
			ui.Flow().
				Class("p").
				StretchItems().
				Spacing(18).
				ItemWidth(360).
				Content(
					newBuiltWithGoappItem().
						Class("fill").
						Image("web/images/prp.png").
						Name("Lofimusic.app").
						Description("App to listen Lo-fi radios.").
						Href("services/prp.md"),
					newBuiltWithGoappItem().
						Class("fill").
						Image("/web/images/dysport.png").
						Name("Dysport Botox").
						Description("Tool to converts Go code into its go/ast representation.").
						Href("services/dysport.md"),
					newBuiltWithGoappItem().
						Class("fill").
						Image("/web/images/filler.png").
						Name("Fillers").
						Description("Face Neck and lips.").
						Href("services/sculptra.md"),
					newBuiltWithGoappItem().
						Class("fill").
						Image("/web/images/microneedling.png").
						Name("Keygean").
						Description("Sign, verify, encrypt and decrypt data with GPG in your browser.").
						Href("services/skinpen.md"),
					newBuiltWithGoappItem().
						Class("fill").
						Image("web/images/iv.png").
						Name("IV Thearapy").
						Description("Hangover, Beauty, Hydration").
						Href("services/ivtherapy.md"),
				),
		)
}

type builtWithGoappItem struct {
	app.Compo

	Iclass       string
	Iimage       string
	Iname        string
	Idescription string
	Ihref        string
}

func newBuiltWithGoappItem() *builtWithGoappItem {
	return &builtWithGoappItem{}
}

func (i *builtWithGoappItem) Class(v string) *builtWithGoappItem {
	i.Iclass = app.AppendClass(i.Iclass, v)
	return i
}

func (i *builtWithGoappItem) Image(v string) *builtWithGoappItem {
	i.Iimage = v
	return i
}

func (i *builtWithGoappItem) Name(v string) *builtWithGoappItem {
	i.Iname = v
	return i
}

func (i *builtWithGoappItem) Description(v string) *builtWithGoappItem {
	i.Idescription = v
	return i
}

func (i *builtWithGoappItem) Href(v string) *builtWithGoappItem {
	i.Ihref = v
	return i
}

func (i *builtWithGoappItem) Render() app.UI {
	return app.A().
		Class(i.Iclass).
		Class("block").
		Class("rounded").
		Class("text-center").
		Class("magnify").
		Class("default").
		Href(i.Ihref).
		Body(
			ui.Block().
				Class("fill").
				Middle().
				Content(
					app.Img().
						Class("hstretch").
						Alt(i.Iname+" tumbnail.").
						Src(i.Iimage),
					app.H3().Text(i.Iname),
					app.Div().
						Class("text-tiny-top").
						Text(i.Idescription),
				),
		)
}
