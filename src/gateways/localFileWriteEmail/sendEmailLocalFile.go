package localFileWriteEmail

import (
	"fmt"
	"github.com/halprin/email-conceal/src/context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var applicationContext = context.ApplicationContext{}

type LocalFileWriteEmailGateway struct{}

func (receiver LocalFileWriteEmailGateway) SendEmail(email []byte, recipients []string) error {
	outputDirectory := os.Args[1]

	filePath := fmt.Sprintf("%s%c%s%s", outputDirectory, filepath.Separator, strings.Join(recipients, ","), "_forwarded.eml")

	err := ioutil.WriteFile(filePath, email, os.ModePerm)

	if err != nil {
		log.Printf("Unable to write out the email %s because %+v", filePath, err)
		return err
	}

	return nil
}
