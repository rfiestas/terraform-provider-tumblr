package tumblr

import (
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/tumblr/tumblr.go"
	"github.com/tumblr/tumblrclient.go"
)

func TestAccPostText_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostTextDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostTextBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_text.first_text", "blog", TestBlog),
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
					resource.TestCheckResourceAttr("tumblr_post_text.first_text", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_text.first_text", "title", "title_first_text"),
				),
			},
			{
				Config: testPostTextUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_text.first_text", "blog", TestBlog),
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
				ExpectError: TestError404NotFound,
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
					resource.TestCheckResourceAttr("tumblr_post_text.first_text", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_text.first_text", "title", "title_first_text"),
				),
			},
			{
				Config:      testPostTextFailure,
				ExpectError: TestError404NotFound,
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

var testPostTextBasic = fmt.Sprintf(`
resource "tumblr_post_text" "first_text" {
  blog  = "%s"
  title = "title_first_text"
  body  = "body_test"
  state = "published"
  tags  = "terraform,terraform provider,tumblr,text,test"
}
`, TestBlog)

var testPostTextUpdate = fmt.Sprintf(`
resource "tumblr_post_text" "first_text" {
	blog  = "%s"
	title = "title_first_text_update"
	body  = "body_test"
	state = "published"
	tags  = "terraform,terraform provider,tumblr,text,test"
}
`, TestBlog)

const testPostTextFailure = `
resource "tumblr_post_text" "first_text" {
	blog  = "NoExistInTumblr"
	title = "title_first_text_failure"
	body  = "body_test"
	state = "published"
	tags  = "terraform,terraform provider,tumblr,text,test"
}
`
