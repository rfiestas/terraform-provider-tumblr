package tumblr

import "regexp"

var TestBlog = "terraform-provider-for-test"

var TestError404NotFound = regexp.MustCompile("404 Not Found")
var TestMustBeAssigned = regexp.MustCompile("must be assigned")
