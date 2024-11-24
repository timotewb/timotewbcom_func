package handler

import (
	"openruntimes/handler/app"

	"github.com/open-runtimes/types-for-go/v4/openruntimes"
)

func Main(Context openruntimes.Context) openruntimes.Response {
	switch Context.Req.Path {
	case "/greeting": return Context.Res.Text(app.Greeting())
	case "/home": return Context.Res.Text(app.Home())
	case "/help": return Context.Res.Text(app.Help())
	default: return Context.Res.Text(app.Fail())
	}
}
