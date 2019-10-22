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

//var TestError404NotFound = regexp.MustCompile("404 Not Found")

func TestAccPostChat_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostChatDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostChatBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_chat.first_chat", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_chat.first_chat", "title", "title_first_chat"),
				),
			},
		},
	})
}
func TestAccPostChat_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostChatDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostChatBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_chat.first_chat", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_chat.first_chat", "title", "title_first_chat"),
				),
			},
			{
				Config: testPostChatUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_chat.first_chat", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_chat.first_chat", "title", "title_first_chat_update"),
				),
			},
		},
	})
}

func TestAccPostChat_WrongCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostChatDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testPostChatFailure,
				ExpectError: TestError404NotFound,
			},
		},
	})
}

func TestAccPostChat_WrongUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostChatDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostChatBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_chat.first_chat", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_chat.first_chat", "title", "title_first_chat"),
				),
			},
			{
				Config:      testPostChatFailure,
				ExpectError: TestError404NotFound,
			},
		},
	})
}

func testAccPostChatDestroy(s *terraform.State) error {
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

var testPostChatBasic = fmt.Sprintf(`
resource "tumblr_post_chat" "first_chat" {
  blog  = "%s"
  title = "title_first_chat"
  conversation  = "foo: ping\nvar: pong"
  state = "published"
  tags  = "terraform,terraform provider,tumblr,chat,test"
}
`, TestBlog)

var testPostChatUpdate = fmt.Sprintf(`
resource "tumblr_post_chat" "first_chat" {
	blog  = "%s"
	title = "title_first_chat_update"
	conversation  = "foo: ping\nvar: pong"
	state = "published"
	tags  = "terraform,terraform provider,tumblr,chat,test"
}
`, TestBlog)

const testPostChatFailure = `
resource "tumblr_post_chat" "first_chat" {
	blog  = "NoExistInTumblr"
	title = "title_first_chat_failure"
	conversation  = "foo: ping\nvar: pong"
	state = "published"
	tags  = "terraform,terraform provider,tumblr,chat,test"
}
`
