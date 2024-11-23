package app

import (
	"fmt"
	"time"
)

func Greeting() (string){
	t := time.Now()
	formattedTime := t.Format("Saturday 3 January 2006 at 03.04PM")

	greeting := `<b>Welcome to the terminal interface!</b>

Use commands to navigate and interact with the website.
To get started, try entering the command <span class='hst-command'>help</span> and hitting enter.

Session start time: %v
    `

	// Use %v format specifier to evaluate expressions
	result := fmt.Sprintf(greeting, formattedTime)

	return result
}