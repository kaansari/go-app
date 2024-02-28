package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
	"gopkg.in/yaml.v2"
)

type MenuItem struct {
	Title string `yaml:"title"`
	URL   string `yaml:"url"`
	Icon  string `yaml:"icon"`
}

type menu struct {
	app.Compo

	Iclass string

	appInstallable bool
	Items          []MenuItem
}

func newMenu() *menu {
	return &menu{}
}

func (m *menu) Class(v string) *menu {
	m.Iclass = app.AppendClass(m.Iclass, v)
	return m
}

func (m *menu) OnNav(ctx app.Context) {
	m.appInstallable = ctx.IsAppInstallable()
}

func (m *menu) OnAppInstallChange(ctx app.Context) {
	m.appInstallable = ctx.IsAppInstallable()
}

func (m *menu) Render() app.UI {

	linkClass := "link heading fit unselectable"

	isFocus := func(path string) string {
		if app.Window().URL().Path == path {
			return "focus"
		}
		return ""
	}

	return ui.Scroll().
		Class("menu").
		Class(m.Iclass).
		HeaderHeight(headerHeight).
		Header(
			ui.Stack().
				Class("fill").
				Middle().
				Content(
					app.Header().Body(
						app.A().
							Class("heading").
							Class("app-title").
							Href("/").
							Text("NooN Spa"),
					),
				),
		).
		Content(
			app.Nav().Body(
				app.Div().Class("separator"),

				ui.Link().
					Class(linkClass).
					Icon(homeSVG).
					Label("Home").
					Href("/").
					Class(isFocus("/")),
				ui.Link().
					Class(linkClass).
					Icon(rocketSVG).
					Label("Services").
					Href("/prp").
					Class(isFocus("/prp")),
				ui.Link().
					Class(linkClass).
					Icon(gridSVG).
					Label("Contact").
					Href("/architecture").
					Class(isFocus("/architecture")),
				ui.Link().
					Class(linkClass).
					Icon(fileTreeSVG).
					Label("Search").
					Href("/search").
					Class(isFocus("/search")),
				ui.Link().
					Class(linkClass).
					Icon(routeSVG).
					Label("Reference").
					Href("/reference").
					Class(isFocus("/reference")),

				app.Div().Class("separator"),

				ui.Link().
					Class(linkClass).
					Icon(twitterSVG).
					Label("Twitter").
					Href(twitterURL),
				ui.Link().
					Class(linkClass).
					Icon(githubSVG).
					Label("GitHub").
					Href(githubURL),
				ui.Link().
					Class(linkClass).
					Icon(opensourceSVG).
					Label("Open Collective").
					Href(openCollectiveURL),

				app.Div().Class("separator"),

				app.If(m.appInstallable,
					ui.Link().
						Class(linkClass).
						Icon(downloadSVG).
						Label("Install").
						OnClick(m.installApp),
				),
				ui.Link().
					Class(linkClass).
					Icon(userLockSVG).
					Label("Privacy Policy").
					Href("/privacy-policy").
					Class(isFocus("/privacy-policy")),

				app.Div().Class("separator"),
			),
		)
}

func (m *menu) loadMenuItems() {
	yamlFile := "data/menu.yaml"
	absPath, err := filepath.Abs(yamlFile)
	if err != nil {
		log.Fatalf("error getting absolute path: %v", err)
	}

	content, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
	}

	err = yaml.Unmarshal(content, &m.Items)
	if err != nil {
		log.Fatalf("error unmarshaling YAML: %v", err)
	}
}

func (m *menu) installApp(ctx app.Context, e app.Event) {
	ctx.NewAction(installApp)
}
