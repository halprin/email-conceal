package external

import (
	"fmt"
	"os"
)

func Cli() {
	applicationContext := CliApplicationContext{}

	concealedEmail := applicationContext.ConcealEmailGateway(os.Args)

	fmt.Println("Concealed e-mail address =", concealedEmail)
}
