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

func TestAccPostPhoto_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostPhotoDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostPhotoBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_photo.first_photo", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_photo.first_photo", "caption", "caption_first_photo"),
				),
			},
		},
	})
}
func TestAccPostPhoto_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostPhotoDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostPhotoBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_photo.first_photo", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_photo.first_photo", "caption", "caption_first_photo"),
				),
			},
			{
				Config: testPostPhotoUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_photo.first_photo", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_photo.first_photo", "caption", "caption_first_photo_update"),
				),
			},
		},
	})
}

func TestAccPostPhoto_WrongCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostPhotoDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testPostPhotoFailure,
				ExpectError: TestError404NotFound,
			},
		},
	})
}

func TestAccPostPhoto_WrongUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostPhotoDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostPhotoBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_photo.first_photo", "blog", TestBlog),
					resource.TestCheckResourceAttr("tumblr_post_photo.first_photo", "caption", "caption_first_photo"),
				),
			},
			{
				Config:      testPostPhotoFailure,
				ExpectError: TestError404NotFound,
			},
		},
	})
}

func TestAccPostPhoto_MissingParameters(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostPhotoDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testPostPhotoMissingParameters,
				ExpectError: TestMustBeAssigned,
			},
		},
	})
}

func testAccPostPhotoDestroy(s *terraform.State) error {
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

var testPostPhotoBasic = fmt.Sprintf(`
resource "tumblr_post_photo" "first_photo" {
	blog    = "%s"
	caption = "caption_first_photo"
	data64  = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z/C/HgAGgwJ/lK3Q6wAAAABJRU5ErkJggg=="
	state   = "published"
	tags    = "terraform,terraform provider,tumblr,photo,test"
}
`, TestBlog)

var testPostPhotoUpdate = fmt.Sprintf(`
resource "tumblr_post_photo" "first_photo" {
	blog    = "%s"
	caption = "caption_first_photo_update"
}
`, TestBlog)

const testPostPhotoFailure = `
resource "tumblr_post_photo" "first_photo" {
	blog    = "NoExistInTumblr"
	caption = "caption_first_photo_failure"
	data64  = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z/C/HgAGgwJ/lK3Q6wAAAABJRU5ErkJggg=="
	state   = "published"
	tags    = "terraform,terraform provider,tumblr,photo,test"
}
`

var testPostPhotoMissingParameters = fmt.Sprintf(`
resource "tumblr_post_photo" "first_photo" {
	blog    = "%s"
	caption = "caption_first_photo_wrong_parameters"
	state   = "published"
	tags    = "terraform,terraform provider,tumblr,photo,test"
}
`, TestBlog)
