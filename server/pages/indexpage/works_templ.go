// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package indexpage

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func works() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h2 class=\"title is-4\">Works</h2><div class=\"columns\"><div class=\"column\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = card("Go", []*CardItem{
			{
				URL:         "https://github.com/syumai/workers",
				Title:       "workers",
				Description: "Go package to run an HTTP server on Cloudflare Workers.",
			},
			{
				URL:         "https://github.com/syumai/protocat",
				Title:       "protocat",
				Description: "A CLI tool to concatenate multiple proto files",
			},
			{
				URL:         "https://github.com/syumai/go-hyperscript",
				Title:       "go-hyperscript",
				Description: "Simple Virtual DOM implemented in Go + WebAssembly",
			},
			{
				URL:         "https://github.com/syumai/syumaigen",
				Title:       "syumaigen",
				Description: "CLI tool to generate syumai's avatar image",
			},
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"column is-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = card("Deno (TypeScript)", []*CardItem{
			{
				URL:         "https://github.com/syumai/dinatra",
				Title:       "dinatra",
				Description: "Small web app framework for Deno",
			},
			{
				URL:         "https://github.com/syumai/dejs",
				Title:       "dejs",
				Description: "EJS engine for Deno",
			},
			{
				URL:         "https://github.com/syumai/dem",
				Title:       "dem",
				Description: "Simple module version manager for Deno",
			},
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"column\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = card("React.js / Vue.js", []*CardItem{
			{
				URL:         "https://github.com/syumai/go-playground-addons",
				Title:       "go-playground-addons",
				Description: "Chrome extension for the Go Playground implemented in React Hooks + TypeScript",
			},
			{
				URL:         "https://github.com/syumai/trollo",
				Title:       "trollo",
				Description: "Simple Trello like kanban app implemented in Vue + Vuex",
			},
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div class=\"columns\"><div class=\"column is-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = card("Ruby", []*CardItem{
			{
				URL:         "https://github.com/syumai/yabass",
				Title:       "yabass",
				Description: "Static site generator",
			},
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"column is-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = card("Others", []*CardItem{
			{
				URL:         "https://gpgsync.herokuapp.com",
				Title:       "gpgsync",
				Description: "Go Playground with coedit mode",
			},
			{
				URL:              "/2020",
				Title:            "2020",
				Description:      "New year content of 2020",
				OpenInCurrentTab: true,
			},
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
