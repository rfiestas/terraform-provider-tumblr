package tumblr

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/tumblr/tumblr.go"
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

func fileBase64(file string) string {
	src, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}
	enc := base64.StdEncoding.EncodeToString(src)
	return enc
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

// Set common fields, some non common fields are not returned by tumblr go client.
func setPostSets(d *schema.ResourceData, res *tumblr.Posts) {
	d.Set("state", res.Get(0).GetSelf().State)
	d.Set("state", res.Get(0).GetSelf().Tags)
	d.Set("date", res.Get(0).GetSelf().Date)
	d.Set("format", res.Get(0).GetSelf().Format)
	d.Set("slug", res.Get(0).GetSelf().Slug)
}
