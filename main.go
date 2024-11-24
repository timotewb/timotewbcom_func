package handler

import (
	"openruntimes/handler/app"

	"github.com/open-runtimes/types-for-go/v4/openruntimes"
)

func Main(Context openruntimes.Context) openruntimes.Response {
	switch Context.Req.Path {
	case "/greeting": return Context.Res.Text(app.Greeting(Context.Req.BodyText()))
	case "/home": return Context.Res.Text(app.Home(Context.Req.BodyText()))
	case "/help": return Context.Res.Text(app.Help(Context.Req.BodyText()))
	default: return Context.Res.Text(app.Fail(Context.Req.Path))
	}
}
