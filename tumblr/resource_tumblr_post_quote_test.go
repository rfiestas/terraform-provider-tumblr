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

func TestAccPostQuote_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostQuoteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostQuoteBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_quote.first_quote", "blog", "terraform-provider-for"),
					resource.TestCheckResourceAttr("tumblr_post_quote.first_quote", "quote", "title_first_quote"),
				),
			},
		},
	})
}
func TestAccPostQuote_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostQuoteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostQuoteBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_quote.first_quote", "blog", "terraform-provider-for"),
					resource.TestCheckResourceAttr("tumblr_post_quote.first_quote", "quote", "title_first_quote"),
				),
			},
			{
				Config: testPostQuoteUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_quote.first_quote", "blog", "terraform-provider-for"),
					resource.TestCheckResourceAttr("tumblr_post_quote.first_quote", "quote", "title_first_quote_update"),
				),
			},
		},
	})
}

func TestAccPostQuote_WrongCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostQuoteDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testPostQuoteFailure,
				ExpectError: error404NotFound,
			},
		},
	})
}

func TestAccPostQuote_WrongUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPostQuoteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testPostQuoteBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tumblr_post_quote.first_quote", "blog", "terraform-provider-for"),
					resource.TestCheckResourceAttr("tumblr_post_quote.first_quote", "quote", "title_first_quote"),
				),
			},
			{
				Config:      testPostQuoteFailure,
				ExpectError: error404NotFound,
			},
		},
	})
}

func testAccPostQuoteDestroy(s *terraform.State) error {
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

const testPostQuoteBasic = `
resource "tumblr_post_quote" "first_quote" {
  blog  = "terraform-provider-for"
  quote = "title_first_quote"
  state = "published"
  tags  = "terraform,terraform provider,tumblr,quote,test"
}
`

const testPostQuoteUpdate = `
resource "tumblr_post_quote" "first_quote" {
	blog  = "terraform-provider-for"
	quote = "title_first_quote_update"
	state = "published"
	tags  = "terraform,terraform provider,tumblr,quote,test"
}
`

const testPostQuoteFailure = `
resource "tumblr_post_quote" "first_quote" {
	blog  = "terraform-provider-tumblr"
	quote = "title_first_quote_failure"
	state = "published"
	tags  = "terraform,terraform provider,tumblr,quote,test"
}
`
