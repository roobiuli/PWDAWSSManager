package Common

import (
	"github.com/roobiuli/PWDKMREP/Password"
	"os"
	"strings"
)

func GetAllVars() []*Password.Passwd {
	var all []*Password.Passwd

	evs := os.Environ()

	for _, v := range evs {
		temps := strings.SplitN(v, "=", 2)
		if temps[2] == "replaceme" {
			evar := &Password.Passwd{
				Name: temps[1],
			}
			all = append(all, evar)
		}
	}
	return all
}


