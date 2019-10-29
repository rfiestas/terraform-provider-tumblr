package tumblr

import (
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/tumblr/tumblr.go"
	"github.com/tumblr/tumblrclient.go"
)

var fieldsChatPosts = []string{"title", "conversation"}

func resourcePostChat() *schema.Resource {
	return &schema.Resource{
		Create: resourcePostChatCreate,
		Read:   resourcePostChatRead,
		Update: resourcePostChatUpdate,
		Delete: resourcePostChatDelete,

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
			"conversation": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["conversation"],
			},
		},
	}
}

func resourcePostChatCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	params := generateParams(d, "chat", append(fieldsAllPosts, fieldsChatPosts...))
	res, err := tumblr.CreatePost(client, d.Get("blog").(string), params)
	if err != nil {
		d.SetId("")
		return err
	}

	d.SetId(uintToString(res.Id))

	return resourcePostChatRead(d, m)
}

func resourcePostChatRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	params := url.Values{}
	params.Add("type", "chat")
	params.Add("id", d.Id())
	res, err := tumblr.GetPosts(client, d.Get("blog").(string), params)
	if err != nil {
		d.SetId("")
		return err
	}

	setPostSets(d, res)

	return nil
}

func resourcePostChatUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	params := generateParams(d, "chat", append(fieldsAllPosts, fieldsChatPosts...))
	err := tumblr.EditPost(client, d.Get("blog").(string), stringToUint(d.Id()), params)
	if err != nil {
		d.SetId("")
		return err
	}

	return resourcePostChatRead(d, m)
}

func resourcePostChatDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	err := tumblr.DeletePost(client, d.Get("blog").(string), stringToUint(d.Id()))
	if err != nil {
		return err
	}

	return nil
}
