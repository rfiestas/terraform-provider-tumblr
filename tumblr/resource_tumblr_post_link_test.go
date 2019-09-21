package tumblr

import (
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/tumblr/tumblr.go"
	"github.com/tumblr/tumblrclient.go"
)

func TestAccPostLink_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostLinkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostLinkBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_link.first_link", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_link.first_link", "title", "title_first_link"),
				),
			},
		},
	})
}
func TestAccPostLink_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostLinkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostLinkBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_link.first_link", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_link.first_link", "title", "title_first_link"),
				),
			},
			{
				Config: testPostLinkUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_link.first_link", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_link.first_link", "title", "title_first_link_update"),
				),
			},
		},
	})
}

func TestAccPostLink_WrongCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostLinkDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testPostLinkFailure,
				ExpectError: TestError404NotFound,
			},
		},
	})
}

func TestAccPostLink_WrongUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostLinkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostLinkBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_link.first_link", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_link.first_link", "title", "title_first_link"),
				),
			},
			{
				Config:      testPostLinkFailure,
				ExpectError: TestError404NotFound,
			},
		},
	})
}

func testAccPostLinkDestroy(s *terraform.State) error {
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

var testPostLinkBasic = fmt.Sprintf(`
resource "tumblr_post_link" "first_link" {
    blog        = "%s"
    title       = "title_first_link"
    description = "description_test"
    url         = "https://terraform-provider-for.tumblr.com/"
    thumbnail   = "https://repository-images.githubusercontent.com/206975372/3a24af00-d275-11e9-8c20-11bd31b5aaaf"
    state       = "published"
    tags        = "terraform,terraform provider,tumblr,link,test"
}
`, TestBlog)

var testPostLinkUpdate = fmt.Sprintf(`
resource "tumblr_post_link" "first_link" {
	blog        = "%s"
	title       = "title_first_link_update"
	description = "description_test"
	url         = "https://terraform-provider-for.tumblr.com/"
	thumbnail   = "https://repository-images.githubusercontent.com/206975372/3a24af00-d275-11e9-8c20-11bd31b5aaaf"
	state       = "published"
	tags        = "terraform,terraform provider,tumblr,link,test"
}
`, TestBlog)

const testPostLinkFailure = `
resource "tumblr_post_link" "first_link" {
	blog        = "NoExistInTumblr"
	title       = "title_first_link_failure"
	description = "description_test"
	url         = "https://terraform-provider-for.tumblr.com/"
	thumbnail   = "https://repository-images.githubusercontent.com/206975372/3a24af00-d275-11e9-8c20-11bd31b5aaaf"
	state       = "published"
	tags        = "terraform,terraform provider,tumblr,link,test"
}
`
