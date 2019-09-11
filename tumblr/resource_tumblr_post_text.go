package tumblr

import (
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tumblr/tumblr.go"
	"github.com/tumblr/tumblrclient.go"
)

var fieldsTextPosts = []string{"title", "body"}

func resourcePostText() *schema.Resource {
	return &schema.Resource{
		Create: resourcePostTextCreate,
		Read:   resourcePostTextRead,
		Update: resourcePostTextUpdate,
		Delete: resourcePostTextDelete,

		Schema: map[string]*schema.Schema{
			"blog":   blogPostSchema(),
			"state":  statePostSchema(),
			"tags":   tagsPostSchema(),
			"date":   datePostSchema(),
			"format": formatPostSchema(),
			"slug":   slugPostSchema(),
			"title": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["title"],
			},
			"body": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["body"],
			},
		},
	}
}

func resourcePostTextCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)
	params := generateParams(d, "text", append(fieldsAllPosts, fieldsTextPosts...))
	res, err := tumblr.CreatePost(client, d.Get("blog").(string), params)
	if err != nil {
		return err
	}

	d.SetId(uintToString(res.Id))

	return resourcePostTextRead(d, m)
}

func resourcePostTextRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*tumblrclient.Client)
	params := url.Values{}
	params.Add("type", "text")
	params.Add("id", d.Id())
	res, err := tumblr.GetPosts(client, d.Get("blog").(string), params)
	if err != nil {
		d.SetId("")
		return nil
	}

	for _, key := range append(fieldsAllPosts, fieldsTextPosts...) {
		value, err := res.Get(0).GetProperty(toCamelCase(key))
		if err == nil {
			d.Set(key, value)
		}
	}

	return nil
}

func resourcePostTextUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)
	params := generateParams(d, "text", append(fieldsAllPosts, fieldsTextPosts...))
	err := tumblr.EditPost(client, d.Get("blog").(string), stringToUint(d.Id()), params)
	if err != nil {
		return err
	}

	return resourcePostTextRead(d, m)
}

func resourcePostTextDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)
	err := tumblr.DeletePost(client, d.Get("blog").(string), stringToUint(d.Id()))
	if err != nil {
		return err
	}

	return resourcePostTextRead(d, m)
}
