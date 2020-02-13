package errors

import (
	"fmt"

	log "github.com/consensys/bpaas-protocols/errors"
	"github.com/juju/errors"
)

type PublishFn func(msg map[string]interface{}) error

func HandleError(errorMessage string, causeError error, message map[string]interface{}, publishFn PublishFn) error {

	if causeError != nil {
		log.LogError(errors.Annotate(causeError, errorMessage))
		errorMessage = fmt.Sprintf("%s: %s", errorMessage, causeError.Error())
	} else {
		log.LogError(errors.New(errorMessage))
	}

	message["error"] = errorMessage
	publishFn(message)

	return errors.New(errorMessage)
}
