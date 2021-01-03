package localFileReader

import (
	"io/ioutil"
)

type LocalFileReader struct {}

func (receiver LocalFileReader) ReadEmail(uri string) ([]byte, error) {
	return ioutil.ReadFile(uri)
}
