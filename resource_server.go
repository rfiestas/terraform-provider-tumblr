package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tumblr/tumblr.go"
	"github.com/tumblr/tumblrclient.go"
)

var fieldsAllPosts = []string{"state", "tags", "date", "format", "slug"}
var fieldsTextPosts = []string{"title", "body"}
var fieldsPhotoPosts = []string{"caption", "link"}

func validateFileExist(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	if _, err := os.Stat(value); os.IsNotExist(err) {
		errs = append(errs, fmt.Errorf("File %s not exist", value))
		return warns, errs
	}
	return warns, errs
}
func validateState(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	var stateList = []string{"private", "draft", "queue", "published"}

	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}

	for _, v := range stateList {
		if v == value {
			return warns, errs
		}
	}
	errs = append(errs, fmt.Errorf("State '%s' is not valid. Choose one of these: %v", value, stateList))
	return warns, errs
}

func transformFileTobase64(fileName string) string {
	f, _ := os.Open(fileName)
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	f.Close()
	encoded := base64.StdEncoding.EncodeToString(content)
	return encoded
}

func transformStringToUint(str string) uint64 {
	u, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return u
}

func resourcePostText() *schema.Resource {
	return &schema.Resource{
		Create: resourcePostTextCreate,
		Read:   resourcePostPhotoRead,
		Update: resourcePostTextUpdate,
		Delete: resourcePostTextDelete,

		Schema: map[string]*schema.Schema{
			"blog": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "blog-identifier",
			},
			"state": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The state of the post. Specify one of the following: published, draft, queue, private",
				ValidateFunc: validateState,
			},
			"tags": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comma-separated tags for this post",
			},
			"date": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The GMT date and time of the post, as a string",
			},
			"format": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Sets the format type of post. Supported formats are: html & markdown",
				Removed:     "Pending to implement, default is html",
			},
			"slug": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Add a short text summary to the end of the post URL",
				Removed:     "Pending to implement",
			},
			"title": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The optional title of the post, HTML entities must be escaped",
			},
			"body": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The full post body, HTML allowed",
			},
		},
	}
}

func resourcePostPhoto() *schema.Resource {
	return &schema.Resource{
		Create: resourcePostPhotoCreate,
		Read:   resourcePostPhotoRead,
		Update: resourcePostPhotoUpdate,
		Delete: resourcePostPhotoDelete,

		Schema: map[string]*schema.Schema{
			"blog": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "blog-identifier",
			},
			"state": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The state of the post. Specify one of the following: published, draft, queue, private",
				ValidateFunc: validateState,
			},
			"tags": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comma-separated tags for this post",
			},
			"date": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The GMT date and time of the post, as a string",
			},
			"format": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Sets the format type of post. Supported formats are: html & markdown",
				Removed:     "Pending to implement, default is html",
			},
			"slug": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Add a short text summary to the end of the post URL",
				Removed:     "Pending to implement",
			},
			"caption": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The user-supplied caption, HTML allowed",
			},
			"link": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The 'click-through URL' for the photo",
			},
			"data64_file": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				Description:  "A file, then the contents of an image file is encoded using base64, limit 10MB",
				ValidateFunc: validateFileExist,
				StateFunc: func(val interface{}) string {
					return filepath.Base(val.(string))
				},
			},
		},
	}
}

func resourcePostPhotoCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	params := generateParams(d, fieldsAllPosts, fieldsPhotoPosts, "photo")
	if d.HasChange("data64_file") {
		params.Add("data64", transformFileTobase64(d.Get("data64_file").(string)))
	}

	res, err := tumblr.CreatePost(client, d.Get("blog").(string), params)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatUint(res.Id, 10))

	return resourcePostPhotoRead(d, m)
}

func resourcePostPhotoRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePostPhotoUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	params := generateParams(d, fieldsAllPosts, fieldsPhotoPosts, "photo")
	if d.HasChange("data64_file") {
		params.Add("data64", transformFileTobase64(d.Get("data64_file").(string)))
	}

	err := tumblr.EditPost(client, d.Get("blog").(string), transformStringToUint(d.Id()), params)
	if err != nil {
		return err
	}
	return resourcePostPhotoRead(d, m)
}

func resourcePostPhotoDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)

	err := tumblr.DeletePost(client, d.Get("blog").(string), transformStringToUint(d.Id()))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourcePostTextCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)
	params := generateParams(d, fieldsAllPosts, fieldsTextPosts, "text")
	res, err := tumblr.CreatePost(client, d.Get("blog").(string), params)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatUint(res.Id, 10))

	return resourcePostPhotoRead(d, m)
}
func resourcePostTextUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)
	params := generateParams(d, fieldsAllPosts, fieldsTextPosts, "text")
	err := tumblr.EditPost(client, d.Get("blog").(string), transformStringToUint(d.Id()), params)
	if err != nil {
		return err
	}

	return resourcePostPhotoRead(d, m)
}

func resourcePostTextDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*tumblrclient.Client)
	err := tumblr.DeletePost(client, d.Get("blog").(string), transformStringToUint(d.Id()))
	if err != nil {
		return err
	}

	return resourcePostPhotoRead(d, m)
}
func generateParams(d *schema.ResourceData, fieldsAll []string, fieldsCustom []string, postType string) url.Values {
	params := url.Values{}
	params.Add("type", postType)
	for _, value := range fieldsAll {
		if d.HasChange(value) {
			params.Add(value, d.Get(value).(string))
		}
	}
	for _, value := range fieldsCustom {
		if d.HasChange(value) {
			params.Add(value, d.Get(value).(string))
		}
	}
	return params
}
