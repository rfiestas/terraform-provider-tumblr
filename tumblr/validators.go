package tumblr

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

func transformFileTobase64(fileName string) string {
	f, _ := os.Open(fileName)
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	f.Close()
	encoded := base64.StdEncoding.EncodeToString(content)
	return encoded
}

func transformStringToUint(str string) uint64 {
	u, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return u
}
