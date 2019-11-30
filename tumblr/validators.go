package tumblr

import (
	"fmt"
	"os"
	"time"
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

func validateDate(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}

	_, err := time.Parse("2006-01-02 15:04:05 MST", value)
	if err != nil {
		errs = append(errs, fmt.Errorf("Date '%s' is not valid format. Format must be '2006-01-02 15:04:05 MST'", value))
		return warns, errs
	}

	return warns, errs
}
func validateData64(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}

	_, err := os.Stat(value)
	if os.IsNotExist(err) {
		errs = append(errs, fmt.Errorf("File '%s' doesn't exist", value))
		return warns, errs
	} else if err != nil {
		errs = append(errs, err)
		return warns, errs
	}

	return warns, errs
}
