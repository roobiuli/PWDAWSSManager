package Common

import (
	"os"
	"strings"

	"github.com/roobiuli/PWDKMREP/Password"
)

func FindVars(s string) []Password.Passwd {
	var allP []Password.Passwd
	allE := os.Environ()

	for _, i := range allE {
		temps := strings.Split(i, "=")
		if temps[1] == s {

			allP = append(allP, Password.Passwd{
				Name: temps[0],
			})
		}
	}
	return allP
}
