package app

import (
	"fmt"
	"strings"
)

func Fail(cmd string) (string){

	resp_string  := `<span class='hst-error'>Error:</span> the <span class='hst-command'>%v</span> command is not recognised.`

	resp_string  = fmt.Sprintf(resp_string, cmd[1:])
	resp_string = strings.ReplaceAll(resp_string , "\n", "<br>")

	return resp_string 
}