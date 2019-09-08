package tumblr

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/tumblr/tumblrclient.go"
)

// Provider returns a terraform.ResourceProvider for the Tumblr provider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"consumer_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONSUMER_KEY", ""),
				Description: descriptions["consumer_key"],
			},
			"consumer_secret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONSUMER_SECRET", ""),
			},
			"user_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("USER_TOKEN", ""),
			},
			"user_token_secret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("USER_TOKEN_SECRET", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"tumblr_post_text":  resourcePostText(),
			"tumblr_post_photo": resourcePostPhoto(),
		},
		ConfigureFunc: providerConfigure,
	}
}

var descriptions map[string]string
var endpointServiceNames []string

func init() {
	descriptions = map[string]string{
		"consumer_key": "PENDING",

		"consumer_secret": "PENDING",

		"user_token": "PENDING",

		"user_token_secret": "PENDING",

		"blog": "blog-identifier",

		"state": "The state of the post. Specify one of the following:\n " +
			" published, draft, queue, private",

		"tags": "Comma-separated tags for this post",

		"data": "The GMT date and time of the post, as a string",

		"format": "Sets the format type of post. Supported formats are:\n " +
			" html & markdown",

		"caption": "The user-supplied caption, HTML allowed",

		"link": "The 'click-through URL' for the photo",

		"data64_file": "A file, then the contents of an image file is encoded using base64,\n " +
			" limit 10MB",

		"title": "The optional title of the post, HTML entities must be escaped",

		"body": "The full post body, HTML allowed",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	consumerKey := d.Get("consumer_key").(string)
	consumerSecret := d.Get("consumer_secret").(string)
	userToken := d.Get("user_token").(string)
	userTokenSecret := d.Get("user_token_secret").(string)

	return tumblrclient.NewClientWithToken(
		consumerKey,
		consumerSecret,
		userToken,
		userTokenSecret,
	), nil
}
