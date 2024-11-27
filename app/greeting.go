package app

import (
	"fmt"
	"strings"
	"time"
)

func Greeting(flags string) (string){
	var resp_string string
	if flags == "" {
		t := time.Now()
		formattedTime := t.Format("Saturday 30 January 2006 at 03.04PM")

		resp_string  = `<b>Welcome to the terminal interface!</b>

Use commands to navigate and interact with the website.
To get started, try entering the command <span class='hst-command'>help</span> and hitting enter.

Session start time: %v`

		resp_string  = fmt.Sprintf(resp_string, formattedTime)
	} else {
		resp_string  = `<span class='hst-error'>Error:</span> the <span class='hst-command'>greeting</span> command does not implement flags '%v'.`
		resp_string  = fmt.Sprintf(resp_string, flags)
	}

	resp_string = strings.ReplaceAll(resp_string , "\n", "<br>")
	return resp_string 
}