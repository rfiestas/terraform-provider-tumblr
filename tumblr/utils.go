package tumblr

import (
	"bufio"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func fileToBase64(filePath string) string {
	f, _ := os.Open(filePath)
	defer f.Close()
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	encoded := base64.StdEncoding.EncodeToString(content)
	return encoded
}

func stringToUint(str string) uint64 {
	u, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return u
}

func uintToString(str uint64) string {
	u := strconv.FormatUint(str, 10)
	return u
}

func hashFileMd5(filePath string) string {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return ""
	}

	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String
}

func generateParams(d *schema.ResourceData, postType string, fields []string) url.Values {
	params := url.Values{}
	params.Add("type", postType)

	for _, value := range fields {
		params.Add(value, d.Get(value).(string))
	}

	return params
}
