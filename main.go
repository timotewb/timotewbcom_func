package handler

import (
	"openruntimes/handler/app"

	"github.com/open-runtimes/types-for-go/v4/openruntimes"
)

type Response struct {
	Message       string `json:"message"`
	Ask       string `json:"ask"`
}

// This Appwrite function will be executed every time your function is triggered
func Main(Context openruntimes.Context) openruntimes.Response {

	switch Context.Req.Path {
	case "/greeting": return Context.Res.Text(app.Greeting())
	case "/home": return Context.Res.Text(app.Home())
	case "/help": return Context.Res.Text(app.Help())
	default: return Context.Res.Text(app.Fail())
	}
	return Context.Res.Json(Response{
		Message:   "Thank you for hitting my API",
		Ask:       "Please done attack me :)",
	})
}
