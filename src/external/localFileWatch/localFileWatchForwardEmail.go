package localFileWatch

import (
	"fmt"
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/controllers/forwardEmail"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var applicationContext = context.ApplicationContext{}

func LocalFileWatcher() {
	directoryToWatch := os.Args[1]
	watchDirectory(directoryToWatch)
}

func watchDirectory(directoryToWatch string) {
	log.Printf("Watching directory %s for e-mail *.dms files", directoryToWatch)

	for {
		directoryContents, err := ioutil.ReadDir(directoryToWatch)
		if err != nil {
			log.Fatalf("Cannot list contents of directory %s", directoryToWatch)
		}

		for _, file := range directoryContents {
			if !fileIsUnforwardedEml(file) {
				continue
			}

			forwardFile(directoryToWatch, file)
			markFileAsForwarded(file.Name())
		}

		time.Sleep(5 * time.Second)
	}
}

var forwardedFiles = make(map[string]interface{})

func markFileAsForwarded(fileName string) {
	forwardedFiles[fileName] = struct{}{}
}

func fileHasBeenForwarded(fileName string) bool {
	_, exists := forwardedFiles[fileName]
	return exists
}

func fileIsUnforwardedEml(file os.FileInfo) bool {
	return !file.IsDir() &&
		!strings.HasSuffix(file.Name(), "_forwarded.eml") &&
		!fileHasBeenForwarded(file.Name()) &&
		strings.HasSuffix(file.Name(), ".eml")
}

func forwardFile(directory string, file os.FileInfo) {
	var forwardEmailController forwardEmail.ForwardEmail
	applicationContext.Resolve(&forwardEmailController)

	filePath := fmt.Sprintf("%s%c%s", directory, filepath.Separator, file.Name())

	log.Printf("Forwarding %s", filePath)
	arguments := map[string]interface{}{
		"url": filePath,
	}
	err := forwardEmailController.ForwardEmail(arguments)

	if err != nil {
		log.Printf("Unable to to forward file %s, error is %+v", filePath, err)
	}
}
