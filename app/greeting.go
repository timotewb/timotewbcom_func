package app

import (
	"fmt"
	"strings"
	"time"
)

func Greeting() (string){
	t := time.Now()
	formattedTime := t.Format("Saturday 3 January 2006 at 03.04PM")

	multiline_string  := `<b>Welcome to the terminal interface!</b>

Use commands to navigate and interact with the website.
To get started, try entering the command <span class='hst-command'>help</span> and hitting enter.

Session start time: %v
    `

	multiline_string = strings.ReplaceAll(multiline_string , "\n", "<br>")
	formatted_string  := fmt.Sprintf(multiline_string, formattedTime)

	return formatted_string 
}