package tumblr

import (
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tumblr/tumblr.go"
	"github.com/tumblr/tumblrclient.go"
)

var fieldsVideoPosts = []string{"caption", "embed", "data"}

func resourcePostVideo() *schema.Resource {
	return &schema.Resource{
		Create: resourcePostVideoCreate,
		Read:   resourcePostVideoRead,
		Update: resourcePostVideoUpdate,
		Delete: resourcePostVideoDelete,

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
			"embed": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   descriptions["embed"],
				ConflictsWith: []string{"data"},
			},
			"data": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   descriptions["data_video"],
				ConflictsWith: []string{"embed"},
				StateFunc: func(val interface{}) string {
					return stringToMd5(val.(string))
				},
				Removed: "Pending to implement, default is external_url",
			},
		},
	}
}

func resourcePostVideoCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	_, embedOk := d.GetOk("embed")
	_, dataOk := d.GetOk("data")

	if !embedOk && !dataOk {
		return fmt.Errorf("One of embed or data must be assigned")
	}

	params := generateParams(d, "video", append(fieldsAllPosts, fieldsVideoPosts...))
	res, err := tumblr.CreatePost(client, d.Get("blog").(string), params)
	if err != nil {
		return err
	}

	d.SetId(uintToString(res.Id))

	return resourcePostVideoRead(d, m)
}

func resourcePostVideoRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	params := url.Values{}
	params.Add("type", "video")
	params.Add("id", d.Id())
	res, err := tumblr.GetPosts(client, d.Get("blog").(string), params)
	if err != nil {
		d.SetId("")
		return nil
	}

	for _, key := range append(fieldsAllPosts, fieldsVideoPosts...) {
		value, err := res.Get(0).GetProperty(toCamelCase(key))
		if err == nil {
			d.Set(key, value)
		}
	}

	return nil
}

func resourcePostVideoUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	params := generateParams(d, "video", append(fieldsAllPosts, fieldsVideoPosts...))
	err := tumblr.EditPost(client, d.Get("blog").(string), stringToUint(d.Id()), params)
	if err != nil {
		return err
	}

	return resourcePostVideoRead(d, m)
}

func resourcePostVideoDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	err := tumblr.DeletePost(client, d.Get("blog").(string), stringToUint(d.Id()))
	if err != nil {
		return err
	}

	return resourcePostVideoRead(d, m)
}
