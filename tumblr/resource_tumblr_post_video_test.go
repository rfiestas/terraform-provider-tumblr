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

func TestAccPostVideo_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostVideoDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostVideoBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_video.first_video", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_video.first_video", "caption", "caption_first_video"),
				),
			},
		},
	})
}
func TestAccPostVideo_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostVideoDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostVideoBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_video.first_video", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_video.first_video", "caption", "caption_first_video"),
				),
			},
			{
				Config: testPostVideoUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_video.first_video", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_video.first_video", "caption", "caption_first_video_update"),
				),
			},
		},
	})
}

func TestAccPostVideo_WrongCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostVideoDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testPostVideoFailure,
				ExpectError: TestError404NotFound,
			},
		},
	})
}

func TestAccPostVideo_WrongUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostVideoDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostVideoBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_video.first_video", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_video.first_video", "caption", "caption_first_video"),
				),
			},
			{
				Config:      testPostVideoFailure,
				ExpectError: TestError404NotFound,
			},
		},
	})
}

func TestAccPostVideo_MissingParameters(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostVideoDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testPostVideoMissingParameters,
				ExpectError: TestMustBeAssigned,
			},
		},
	})
}

func testAccPostVideoDestroy(s *terraform.State) error {
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

var testPostVideoBasic = fmt.Sprintf(`
resource "tumblr_post_video" "first_video" {
	blog    = "%s"
	caption = "caption_first_video"
	embed   = "https://www.youtube.com/watch?v=TMayVLSQ6yM"
	state   = "published"
	tags    = "terraform,terraform provider,tumblr,video,test"
}
`, TestBlog)

var testPostVideoUpdate = fmt.Sprintf(`
resource "tumblr_post_video" "first_video" {
	blog    = "%s"
	caption = "caption_first_video_update"
}
`, TestBlog)

const testPostVideoFailure = `
resource "tumblr_post_video" "first_video" {
	blog    = "NoExistInTumblr"
	caption = "caption_first_video_failure"
	embed   = "https://www.youtube.com/watch?v=TMayVLSQ6yM"
	state   = "published"
	tags    = "terraform,terraform provider,tumblr,video,test"
}
`

var testPostVideoMissingParameters = fmt.Sprintf(`
resource "tumblr_post_video" "first_video" {
	blog    = "%s"
	caption = "caption_first_video_wrong_parameters"
	state   = "published"
	tags    = "terraform,terraform provider,tumblr,video,test"
}
`, TestBlog)
