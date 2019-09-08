package tumblr

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tumblr/tumblr.go"
	"github.com/tumblr/tumblrclient.go"
)

var fieldsQuotePosts = []string{"quote", "source"}

func resourcePostQuote() *schema.Resource {
	return &schema.Resource{
		Create: resourcePostQuoteCreate,
		Read:   resourcePostQuoteRead,
		Update: resourcePostQuoteUpdate,
		Delete: resourcePostQuoteDelete,

		Schema: map[string]*schema.Schema{
			"blog":   blogPostSchema(),
			"state":  statePostSchema(),
			"tags":   tagsPostSchema(),
			"date":   datePostSchema(),
			"format": formatPostSchema(),
			"slug":   slugPostSchema(),
			"quote": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["quote"],
			},
			"source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["source"],
			},
		},
	}
}

func resourcePostQuoteCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)
	params := generateParams(d, "quote", append(fieldsAllPosts, fieldsQuotePosts...))
	res, err := tumblr.CreatePost(client, d.Get("blog").(string), params)
	if err != nil {
		return err
	}

	d.SetId(uintToString(res.Id))

	return resourcePostQuoteRead(d, m)
}

func resourcePostQuoteRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePostQuoteUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)
	params := generateParams(d, "quote", append(fieldsAllPosts, fieldsQuotePosts...))
	err := tumblr.EditPost(client, d.Get("blog").(string), stringToUint(d.Id()), params)
	if err != nil {
		return err
	}

	return resourcePostQuoteRead(d, m)
}

func resourcePostQuoteDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)
	err := tumblr.DeletePost(client, d.Get("blog").(string), stringToUint(d.Id()))
	if err != nil {
		return err
	}

	return resourcePostQuoteRead(d, m)
}
