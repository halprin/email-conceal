package gateways

import (
	"github.com/halprin/email-conceal/forwarder/context"
	"io/ioutil"
)

func FileReadEmailGateway(url string, applicationContext context.ApplicationContext) ([]byte, error) {
	return ioutil.ReadFile(url)
}
