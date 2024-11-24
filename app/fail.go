package app

import (
	"strings"
)

func Fail() (string){

	multiline_string  := `<span class='hst-error'>Error:</span> the <span class='hst-command'>help</span> command not recognised.`

	multiline_string = strings.ReplaceAll(multiline_string , "\n", "<br>")

	return multiline_string 
}