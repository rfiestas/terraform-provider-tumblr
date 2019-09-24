package tumblr

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

var camelCase = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")

func stringToUint(str string) uint64 {
	u, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return u
}

func uintToString(integer uint64) string {
	u := strconv.FormatUint(integer, 10)
	return u
}

func stringToMd5(str string) string {
	var returnMD5String string

	hasher := md5.New()
	hasher.Write([]byte(str))
	returnMD5String = hex.EncodeToString(hasher.Sum(nil))

	return returnMD5String
}

// toCamelCase foo_var_foo_var to FooVarFooVar
func toCamelCase(str string) string {
	return camelCase.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}

func generateParams(d *schema.ResourceData, postType string, fields []string) url.Values {
	params := url.Values{}
	params.Add("type", postType)

	for _, value := range fields {
		if d.HasChange(value) {
			params.Add(value, d.Get(value).(string))
		}
	}

	return params
}
