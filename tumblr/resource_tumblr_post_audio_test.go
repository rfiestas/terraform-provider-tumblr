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

func TestAccPostAudio_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostAudioDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostAudioBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_audio.first_audio", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_audio.first_audio", "caption", "caption_first_audio"),
				),
			},
		},
	})
}
func TestAccPostAudio_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostAudioDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostAudioBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_audio.first_audio", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_audio.first_audio", "caption", "caption_first_audio"),
				),
			},
			{
				Config: testPostAudioUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_audio.first_audio", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_audio.first_audio", "caption", "caption_first_audio_update"),
				),
			},
		},
	})
}

func TestAccPostAudio_WrongCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostAudioDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testPostAudioFailure,
				ExpectError: error404NotFound,
			},
		},
	})
}

func TestAccPostAudio_WrongUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostAudioDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostAudioBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_audio.first_audio", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_audio.first_audio", "caption", "caption_first_audio"),
				),
			},
			{
				Config:      testPostAudioFailure,
				ExpectError: error404NotFound,
			},
		},
	})
}

func TestAccPostAudio_MissingParameters(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostAudioDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testPostAudioMissingParameters,
				ExpectError: mustBeAssigned,
			},
		},
	})
}

func testAccPostAudioDestroy(s *terraform.State) error {
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

var testPostAudioBasic = fmt.Sprintf(`
resource "tumblr_post_audio" "first_audio" {
	blog         = "%s"
	caption      = "caption_first_audio"
	external_url = "https://soundcloud.com/club-bizarre/free-download-kraftwerk-we-are"
	state        = "published"
	tags         = "terraform,terraform provider,tumblr,audio,test"
}
`, TestBlog)

var testPostAudioUpdate = fmt.Sprintf(`
resource "tumblr_post_audio" "first_audio" {
	blog    = "%s"
	caption = "caption_first_audio_update"
}
`, TestBlog)

const testPostAudioFailure = `
resource "tumblr_post_audio" "first_audio" {
	blog         = "NoExistInTumblr"
	caption      = "caption_first_audio_failure"
	external_url = "https://soundcloud.com/club-bizarre/free-download-kraftwerk-we-are"
	state        = "published"
	tags         = "terraform,terraform provider,tumblr,audio,test"
}
`

const testPostAudioMissingParameters = `
resource "tumblr_post_audio" "first_audio" {
	blog    = "NoExistInTumblr"
	caption = "caption_first_audio_wrong_parameters"
	state   = "published"
	tags    = "terraform,terraform provider,tumblr,audio,test"
}
`
