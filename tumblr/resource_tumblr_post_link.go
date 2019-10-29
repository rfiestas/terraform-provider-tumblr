package tumblr

import (
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/tumblr/tumblr.go"
	"github.com/tumblr/tumblrclient.go"
)

var fieldsLinkPosts = []string{"title", "url", "description", "thumbnail", "excerpt", "author"}

func resourcePostLink() *schema.Resource {
	return &schema.Resource{
		Create: resourcePostLinkCreate,
		Read:   resourcePostLinkRead,
		Update: resourcePostLinkUpdate,
		Delete: resourcePostLinkDelete,

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
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["url"],
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["description"],
			},
			"thumbnail": {
				Type:        schema.TypeString,
				Required:    true, //Optional on Tumblr api documentation.
				Description: descriptions["thumbnail"],
			},
			"excerpt": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["excerpt"],
			},
			"author": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["author"],
			},
		},
	}
}

func resourcePostLinkCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	params := generateParams(d, "link", append(fieldsAllPosts, fieldsLinkPosts...))
	res, err := tumblr.CreatePost(client, d.Get("blog").(string), params)
	if err != nil {
		d.SetId("")
		return err
	}

	d.SetId(uintToString(res.Id))

	return resourcePostLinkRead(d, m)
}

func resourcePostLinkRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	params := url.Values{}
	params.Add("type", "link")
	params.Add("id", d.Id())
	res, err := tumblr.GetPosts(client, d.Get("blog").(string), params)
	if err != nil {
		d.SetId("")
		return err
	}

	setPostSets(d, res)

	return nil
}

func resourcePostLinkUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	params := generateParams(d, "link", append(fieldsAllPosts, fieldsLinkPosts...))
	err := tumblr.EditPost(client, d.Get("blog").(string), stringToUint(d.Id()), params)
	if err != nil {
		d.SetId("")
		return err
	}

	return resourcePostLinkRead(d, m)
}

func resourcePostLinkDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	err := tumblr.DeletePost(client, d.Get("blog").(string), stringToUint(d.Id()))
	if err != nil {
		return err
	}

	return nil
}
