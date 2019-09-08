package tumblr

import (
	"fmt"
	"os"
)

func validateFileExist(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	if _, err := os.Stat(value); os.IsNotExist(err) {
		errs = append(errs, fmt.Errorf("File %s not exist", value))
		return warns, errs
	}
	return warns, errs
}

func validateState(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	var stateList = []string{"private", "draft", "queue", "published"}

	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}

	for _, v := range stateList {
		if v == value {
			return warns, errs
		}
	}
	errs = append(errs, fmt.Errorf("State '%s' is not valid. Choose one of these: %v", value, stateList))
	return warns, errs
}
