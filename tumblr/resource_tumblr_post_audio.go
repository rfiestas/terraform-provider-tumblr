package tumblr

import (
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tumblr/tumblr.go"
	"github.com/tumblr/tumblrclient.go"
)

var fieldsAudioPosts = []string{"caption", "external_url"}

func resourcePostAudio() *schema.Resource {
	return &schema.Resource{
		Create: resourcePostAudioCreate,
		Read:   resourcePostAudioRead,
		Update: resourcePostAudioUpdate,
		Delete: resourcePostAudioDelete,

		Schema: map[string]*schema.Schema{
			"blog":   blogPostSchema(),
			"state":  statePostSchema(),
			"tags":   tagsPostSchema(),
			"date":   datePostSchema(),
			"format": formatPostSchema(),
			"slug":   slugPostSchema(),
			"caption": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["caption"],
			},
			"external_url": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   descriptions["external_url"],
				ConflictsWith: []string{"data"},
			},
			"data": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   descriptions["data_audio"],
				ConflictsWith: []string{"external_url"},
				StateFunc: func(val interface{}) string {
					return stringToMd5(val.(string))
				},
				Removed: "Pending to implement, default is external_url",
			},
		},
	}
}

func resourcePostAudioCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)
	_, externalURLOk := d.GetOk("external_url")
	_, dataOk := d.GetOk("data")

	if !externalURLOk && !dataOk {
		return fmt.Errorf("One of external_url or data must be assigned")
	}

	params := generateParams(d, "audio", append(fieldsAllPosts, fieldsAudioPosts...))
	res, err := tumblr.CreatePost(client, d.Get("blog").(string), params)
	if err != nil {
		return err
	}

	d.SetId(uintToString(res.Id))

	return resourcePostAudioRead(d, m)
}

func resourcePostAudioRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	_, externalURLOk := d.GetOk("external_url")
	_, dataOk := d.GetOk("data")

	if !externalURLOk && !dataOk {
		return fmt.Errorf("One of external_url or data must be assigned")
	}

	params := url.Values{}
	params.Add("type", "audio")
	params.Add("id", d.Id())
	res, err := tumblr.GetPosts(client, d.Get("blog").(string), params)
	if err != nil {
		d.SetId("")
		return nil
	}

	for _, key := range append(fieldsAllPosts, fieldsAudioPosts...) {
		value, err := res.Get(0).GetProperty(toCamelCase(key))
		if err == nil {
			d.Set(key, value)
		}
	}

	return nil
}

func resourcePostAudioUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)
	params := generateParams(d, "audio", append(fieldsAllPosts, fieldsAudioPosts...))
	err := tumblr.EditPost(client, d.Get("blog").(string), stringToUint(d.Id()), params)
	if err != nil {
		return err
	}

	return resourcePostAudioRead(d, m)
}

func resourcePostAudioDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)
	err := tumblr.DeletePost(client, d.Get("blog").(string), stringToUint(d.Id()))
	if err != nil {
		return err
	}

	return resourcePostAudioRead(d, m)
}
