package tumblr

import (
	"fmt"
)

func validateState(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	var stateList = []string{"private", "draft", "queue", "published"}

	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}

	for _, item := range stateList {
		if item == value {
			return warns, errs
		}
	}
	errs = append(errs, fmt.Errorf("State '%s' is not valid. Choose one of these: %v", value, stateList))

	return warns, errs
}
