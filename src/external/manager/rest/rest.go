package rest

import (
	"github.com/halprin/email-conceal/src/external/manager/localRest"
)

func Rest() {
	//just use local rest version since it is the same for now
	localRest.Rest()
}
