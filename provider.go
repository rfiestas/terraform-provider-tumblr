package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tumblr/tumblrclient.go"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"consumer_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONSUMER_KEY", ""),
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
