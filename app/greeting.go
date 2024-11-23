package app

import (
	"fmt"
	"time"
)

func Greeting() (string){
    t := time.Now()
    formattedTime := t.Format("%A %d %B %Y at %I.%M%p")

	greeting := `<b>Welcome to the terminal interface!</b>

    Use commands to navigate and interact with the website.
    To get started, try entering the command <span class='hst-command'>help</span> and hitting enter.

    Session start time: {{formattedTime}}
    `
    result := fmt.Sprintf(greeting, map[string]interface{}{
        "time": formattedTime,
    })

	return result
}