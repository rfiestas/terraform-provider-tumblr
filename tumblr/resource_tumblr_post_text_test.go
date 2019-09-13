package tumblr

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/tumblr/tumblr.go"
	"github.com/tumblr/tumblrclient.go"
)

var error404NotFound = regexp.MustCompile("404 Not Found")

func TestAccPostText_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostTextDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostTextBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_text.first_text", "blog", "terraform-provider-for"),
					resource.TestCheckResourceAttr("tumblr_post_text.first_text", "title", "title_first_text"),
				),
			},
		},
	})
}
func TestAccPostText_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostTextDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostTextBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_text.first_text", "blog", "terraform-provider-for"),
					resource.TestCheckResourceAttr("tumblr_post_text.first_text", "title", "title_first_text"),
				),
			},
			{
				Config: testPostTextUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_text.first_text", "blog", "terraform-provider-for"),
					resource.TestCheckResourceAttr("tumblr_post_text.first_text", "title", "title_first_text_update"),
				),
			},
		},
	})
}

func TestAccPostText_WrongCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostTextDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testPostTextFailure,
				ExpectError: error404NotFound,
			},
		},
	})
}

func TestAccPostText_WrongUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostTextDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostTextBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_text.first_text", "blog", "terraform-provider-for"),
					resource.TestCheckResourceAttr("tumblr_post_text.first_text", "title", "title_first_text"),
				),
			},
			{
				Config:      testPostTextFailure,
				ExpectError: error404NotFound,
			},
		},
	})
}

func testAccPostTextDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*tumblrclient.Client)
	for _, r := range s.RootModule().Resources {
		params := url.Values{}
		params.Add("type", r.Primary.Attributes["type"])
		params.Add("id", r.Primary.ID)
		_, err := tumblr.GetPosts(client, r.Primary.Attributes["blog"], params)
		if err != nil {
			if strings.Contains(err.Error(), "404 Not Found") {
				continue
			}
			return fmt.Errorf("Received an error retrieving post %s", err)
		}
		return fmt.Errorf("Post still exists")
	}
	return nil
}

const testPostTextBasic = `
resource "tumblr_post_text" "first_text" {
  blog  = "terraform-provider-for"
  title = "title_first_text"
  body  = "body_test"
  state = "published"
  tags  = "terraform,terraform provider,tumblr,text,test"
}
`

const testPostTextUpdate = `
resource "tumblr_post_text" "first_text" {
	blog  = "terraform-provider-for"
	title = "title_first_text_update"
	body  = "body_test"
	state = "published"
	tags  = "terraform,terraform provider,tumblr,text,test"
}
`

const testPostTextFailure = `
resource "tumblr_post_text" "first_text" {
	blog  = "terraform-provider-tumblr"
	title = "title_first_text_failure"
	body  = "body_test"
	state = "published"
	tags  = "terraform,terraform provider,tumblr,text,test"
}
`
