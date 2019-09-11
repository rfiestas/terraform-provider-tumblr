package tumblr

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tumblr/tumblr.go"
	"github.com/tumblr/tumblrclient.go"
)

var fieldsPhotoPosts = []string{"caption", "link", "data`"}

func resourcePostPhoto() *schema.Resource {
	return &schema.Resource{
		Create: resourcePostPhotoCreate,
		Read:   resourcePostPhotoRead,
		Update: resourcePostPhotoUpdate,
		Delete: resourcePostPhotoDelete,

		Schema: map[string]*schema.Schema{
			"blog":   blogPostSchema(),
			"state":  statePostSchema(),
			"tags":   tagsPostSchema(),
			"date":   datePostSchema(),
			"format": formatPostSchema(),
			"slug":   slugPostSchema(),
			"caption": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["caption"],
			},
			"link": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["link"],
			},
			"source": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["source_photo"],
				Removed:     "Pending to implement, default is data64",
			},
			"data": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["data"],
				Removed:     "Pending to implement, default is data64",
			},
			"data64": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["data64"],
				StateFunc: func(val interface{}) string {
					return stringToMd5(val.(string))
				},
			},
		},
	}
}

func resourcePostPhotoCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)
	params := generateParams(d, "photo", append(fieldsAllPosts, fieldsPhotoPosts...))
	//if d.HasChange("data64") {
	//	params.Add("data64", fileToBase64(d.Get("data64_file").(string)))
	//}

	res, err := tumblr.CreatePost(client, d.Get("blog").(string), params)
	if err != nil {
		return err
	}
	d.SetId(uintToString(res.Id))

	return resourcePostPhotoRead(d, m)
}

func resourcePostPhotoRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePostPhotoUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	params := generateParams(d, "photo", append(fieldsAllPosts, fieldsPhotoPosts...))
	//if d.HasChange("data64_file") {
	//	params.Add("data64", fileToBase64(d.Get("data64_file").(string)))
	//}

	err := tumblr.EditPost(client, d.Get("blog").(string), stringToUint(d.Id()), params)
	if err != nil {
		return err
	}
	return resourcePostPhotoRead(d, m)
}

func resourcePostPhotoDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	err := tumblr.DeletePost(client, d.Get("blog").(string), stringToUint(d.Id()))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
